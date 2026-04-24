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

func (r *ProductRepository) GetAll(page, limit int, search string) ([]models.Product, error) {

	offset := (page - 1) * limit

	query := `
	SELECT id,name,description,price,stock
	FROM products
	WHERE LOWER(name) LIKE LOWER($1)
	ORDER BY id DESC
	LIMIT $2 OFFSET $3
	`

	rows, err := r.DB.Query(query, "%"+search+"%", limit, offset)
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

func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	query := `
		SELECT id, name, description, price, stock
		FROM products
		WHERE id = $1
	`
	var p models.Product

	err := r.DB.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Stock,
	)

	if err != nil {	
		return nil, err
	}

	return &p, nil

}

func (r *ProductRepository) Update(id string, p *models.Product) error {
	query := `
		UPDATE products
		SET name=$1, description=$2, price=$3, stock=$4
		WHERE id=$5
	`

	_, err := r.DB.Exec(query, p.Name, p.Description, p.Price, p.Stock, id)

	return err

}

func (r *ProductRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
