package config

import (
	"database/sql"
	"log"
)

func RunMigrations(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	productQuery := `
		CREATE TABLE IF NOT EXISTS products(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		price NUMERIC(10,2) NOT NULL,
		stock INT NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
		);	
	`
	_, err = db.Exec(productQuery)
	if err != nil {
		log.Fatal("products migration failed:", err)
	}

	cartQuery := `
	CREATE TABLE IF NOT EXISTS cart_items (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
	quantity INT NOT NULL CHECK (quantity > 0),
	created_at TIMESTAMP DEFAULT NOW(),
	UNIQUE(user_id, product_id)
);`

	orderQuery := `
	CREATE TABLE IF NOT EXISTS orders (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
	total NUMERIC(10,2) NOT NULL,
	status TEXT NOT NULL DEFAULT 'placed',
	created_at TIMESTAMP DEFAULT NOW()
);`

	orderItemQuery := `
	CREATE TABLE IF NOT EXISTS order_items (
	id SERIAL PRIMARY KEY,
	order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
	product_id INT NOT NULL REFERENCES products(id),
	product_name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL,
	quantity INT NOT NULL
);`

	_, err = db.Exec(cartQuery)
	if err != nil {
		log.Fatal("cart migration failed:", err)
	}

	_, err = db.Exec(orderQuery)
	if err != nil {
		log.Fatal("orders migration failed:", err)
	}

	_, err = db.Exec(orderItemQuery)
	if err != nil {
		log.Fatal("order items migration failed:", err)
	}

	log.Println("Migrations completed")
}
