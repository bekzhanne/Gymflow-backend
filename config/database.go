package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
//глобальная переменная через которую весь проект получает доступ к базе
var DB *gorm.DB


func ConnectDatabase() {
	//формируем строку для подключения к постгре
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	) //зачем это все? чтобы не держать пароль  в коде

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // подключаем горм
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = database // ну и сохраняем
	fmt.Println("Connected to PostgreSQL successfully!")
}