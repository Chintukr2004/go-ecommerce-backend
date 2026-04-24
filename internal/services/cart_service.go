package services

import (
	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/repository"
)

type CartService struct {
	Repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{
		Repo: repo,
	}
}

func (s *CartService) Add(userID int, productID int, qty int) error {
	return s.Repo.Add(userID, productID, qty)
}

func (s *CartService) GetByUser(userID int) ([]models.CartItem, error) {
	return s.Repo.GetByUser(userID)
}
