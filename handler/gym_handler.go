package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gymflow/models"
	"gymflow/service"
)
//хэндлер хранит ссылку на сервис
//через сервис он вызывает бизнес логику
//хэндлер не знает про репозиторий и БД
type GymHandler struct {
	service *service.GymService
}

func NewGymHandler(service *service.GymService) *GymHandler {
	return &GymHandler{service: service}
}

// создаем новый зал(только админ) 
// c *gin.Context - это объект от Gin и он содержит:
// хттп запрос, параметры урл, методы для ответа (джисонка и статус)
func (h *GymHandler) CreateGym(c *gin.Context) {
	var gym models.Gym //создаем пустой объект

	// Парсим JSON из запроса(body) и заполняем
	if err := c.ShouldBindJSON(&gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызываем Service
	if err := h.service.CreateGym(&gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Успешный ответ
	c.JSON(http.StatusCreated, gin.H{
		"message": "gym created successfully",
		"gym":     gym,
	})
}

// ListGyms - получить список всех залов
func (h *GymHandler) ListGyms(c *gin.Context) {
	gyms, err := h.service.ListGyms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load gyms"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"gyms": gyms,
	})
}

// GetGym - получить зал по ID
func (h *GymHandler) GetGym(c *gin.Context) {
	// Получаем ID из URL параметра
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid gym id"})
		return
	}

	gym, err := h.service.GetGymByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "gym not found"})
		return
	}

	c.JSON(http.StatusOK, gym)
}