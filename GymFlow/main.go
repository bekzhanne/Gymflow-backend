package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gymflow/auth"
	"gymflow/config"
	"gymflow/handler"
	"gymflow/models"
	"gymflow/repository"
	"gymflow/service"
)

func main() {
	cfg := config.Load()

	// --- DB init ---
	db, err := gorm.Open(postgres.Open(cfg.DB_DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// миграция таблицы пользователей
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// --- DI ---
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTExpireMins)
	userHandler := handler.NewUserHandler(userService, jwtService)

	r := gin.Default()

	// публичные эндпоинты
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
	}

	// защищённые эндпоинты
	api := r.Group("/api")
	api.Use(auth.JWTAuthMiddleware(jwtService))
	{
		api.GET("/me", userHandler.Me)
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on %s ...", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
