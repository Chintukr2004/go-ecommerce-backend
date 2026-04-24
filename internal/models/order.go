package models

type Order struct {
	ID     int     `json:"id"`
	Total  float64 `json:"total"`
	Status string  `json:"status"`
}

type OrderItem struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
