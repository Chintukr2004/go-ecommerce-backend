package services

import (
	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		Repo: repo,
	}
}

func (s *ProductService) Create(p *models.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) GetAll(page, limit int, search string) ([]models.Product, error) {
	return s.Repo.GetAll(page, limit, search)
}

func (s *ProductService) GetByID(id string) (*models.Product, error) {
	return s.Repo.GetByID(id)
}

func (s *ProductService) Update(id string, p *models.Product) error {
	return s.Repo.Update(id, p)
}

func (s *ProductService) Delete(id string) error {
	return s.Repo.Delete(id)
}
