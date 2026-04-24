package services

import (
	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{
		Repo: repo,
	}
}

func (s *OrderService) Checkout(userID int) error {
	return s.Repo.Checkout(userID)
}

func (s *OrderService) GetByUser(userID int) ([]models.Order, error) {
	return s.Repo.GetByUser(userID)
}
