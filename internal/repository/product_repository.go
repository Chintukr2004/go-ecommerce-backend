package repository

import (
	"database/sql"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) Create(p *models.Product) error {
	query := `
		INSERT INTO products (name, description, price, stock)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(query, p.Name, p.Description, p.Price, p.Stock)
	return err
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := r.DB.Query(`
	SELECT id,name,description,price,stock
	FROM products
	ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product

		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
