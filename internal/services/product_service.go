package services

import (
	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
	"github.com/Chintukr2004/go-ecommerce-backend/internal/repository"
)

type ProductService struct{
	Repo *repository.ProductRepository
}

func NewProductService( repo *repository.ProductRepository) *ProductService{
	return &ProductService{
		Repo: repo,
	}
}


func( s *ProductService) Create(p *models.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) GetAll()([]models.Product, error){
	return s.Repo.GetAll()
}