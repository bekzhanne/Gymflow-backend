package repository

import (
	"gymflow/models"

	"gorm.io/gorm"
)

type GymRepository struct {
	db *gorm.DB
}

func NewGymRepository(db *gorm.DB) *GymRepository {
	return &GymRepository{db: db}
}

// Создаем новый зал
func (r *GymRepository) Create(gym *models.Gym) error {
	return r.db.Create(gym).Error
}

// выводим весь список залов
func (r *GymRepository) GetAll() ([]models.Gym, error) {
	var gyms []models.Gym
	err := r.db.Find(&gyms).Error
	return gyms, err
}

// ищем зал по айди
func (r *GymRepository) GetByID(id uint) (*models.Gym, error) {
	var gym models.Gym
	err := r.db.First(&gym, id).Error
	if err != nil {
		return nil, err
	}
	return &gym, nil
}