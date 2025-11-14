package config

import (
	"log"
	"gymflow/models"
)

func RunMigrations() {
	err := DB.AutoMigrate( // горм создает таблицы
		&models.User{},
		&models.Gym{},
		&models.Booking{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully!")
}
