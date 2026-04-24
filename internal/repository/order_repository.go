package repository

import (
	"database/sql"
	"errors"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderReposiotry(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}

}

func (r *OrderRepository) Checkout(userID int) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rows, err := tx.Query(`
		SELECT p.id, p.name, p.price, p.stock, c.quantity
		FROM cart_items c
		JOIN products p ON p.id = c.product_id
		WHERE c.user_id = $1
	`, userID)

	if err != nil {
		return err
	}

	defer rows.Close()

	type item struct {
		ProductID int
		Name      string
		Price     float64
		Stock     int
		Qty       int
	}

	var items []item
	var total float64

	for rows.Next() {
		var i item

		err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Price,
			&i.Stock,
			&i.Qty,
		)
		if err != nil {
			return err
		}
		if i.Qty > i.Stock {
			return errors.New("insufficient stock")
		}

		total += i.Price * float64(i.Qty)
		items = append(items, i)
	}

	if len(items) == 0 {
		return errors.New("cart is empty")
	}

	var orderID int

	err = tx.QueryRow(`
	INSERT INTO orders (user_id, total, status)
	VALUES ($1, $2, 'placed')
	RETURNING id
	`, userID, total).Scan(&orderID)

	if err != nil {
		return err
	}

	//insert items+reduce stock + clear cart

	for _, i := range items {
		_, err = tx.Exec(`
		INSERT INTO order_items (order_id, product_id, product_name, price, quantity)
		VALUES ($1, $2, $3, $4, $5)`, orderID, i.ProductID, i.Name, i.Price, i.Qty)

		if err != nil {
			return err
		}

		_, err := tx.Exec(`
			UPDATE products 
			SET stock = stock - $1	
			WHERE id  = $2
		`, i.Qty, i.ProductID)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(`
		DELETE FROM cart_items 
		WHERE user_id = $1
	`, userID)

	if err != nil {
		return err
	}

	return tx.Commit()

}

//order history

func (r *OrderRepository) GetByUser(userID int) ([]models.Order, error) {
	rows, err := r.DB.Query(`
		SELECT id, total, status
		FROM orders
		WHERE user_id = $1
		ORDER BY id DESC
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var o models.Order

		err := rows.Scan(&o.ID, &o.Total, &o.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)

	}

	return orders, nil
}
