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

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
