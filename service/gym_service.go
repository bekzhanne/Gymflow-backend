package service

import (
	"errors"
	"gymflow/models"
	"gymflow/repository"
)

type GymService struct {
	repo *repository.GymRepository
}

// создаем новый сервис и он принимает только репозиторий(внедрение зависимостей)
func NewGymService(repo *repository.GymRepository) *GymService {
	return &GymService{repo: repo}
}

// создаем зал
func (s *GymService) CreateGym(gym *models.Gym) error {
	// Валидация
	if gym.Name == "" {
		return errors.New("gym name is required")
	}
	if gym.Capacity < 0 {
		return errors.New("capacity cannot be negative")
	}
	
	// Вызываем репозиторий
	return s.repo.Create(gym)
}

// ListGyms - получить все залы
func (s *GymService) ListGyms() ([]models.Gym, error) {
	return s.repo.GetAll()
}

// GetGymByID - получить зал по ID
func (s *GymService) GetGymByID(id uint) (*models.Gym, error) {
	return s.repo.GetByID(id)
}