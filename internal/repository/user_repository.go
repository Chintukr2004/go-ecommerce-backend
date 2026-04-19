package repository

import (
	"database/sql"

	"github.com/Chintukr2004/go-ecommerce-backend/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name, email, password)
		VALUES($1, $2, $3)
	`
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}
