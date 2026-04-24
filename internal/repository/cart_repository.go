package repository

import (
	"database/sql"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{
		DB: db,
	}
}

func (r *CartRepository) Add(userID int, productID int, qty int) error {
	query := `
	INSERT INTO cart_items (user_id, product_id, quantity)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id, product_id)
	DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity
	`
	_, err := r.DB.Exec(query, userID, productID, qty)
	return err
}

func (r *CartRepository) GetByUser(userID int) ([]models.CartItem, error) {
	query := `
	SELECT c.id, p.id, p.name, p.price, c.quantity
	FROM cart_items c
	JOIN prodcuts p ON p.id = c.product_id
	WHERE c.user_id = $1
	ORDER BY c.id DESC
	`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem

	for rows.Next() {
		var item models.CartItem
		err := rows.Scan(
			&item.ID,
			&item.ProductID,
			&item.ProductName,
			&item.Price,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)

	}

	return items, nil
}
