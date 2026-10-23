package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rahulreghupdm/tashub-project/internal/database"
	"github.com/rahulreghupdm/tashub-project/internal/models"
)

func CreateUser(email, hashedPassword string) error {
	_, err := database.DB.Exec(
		context.Background(),
		"INSERT INTO users (email, password) VALUES ($1, $2)",
		email, hashedPassword,
	)
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.QueryRow(
		context.Background(),
		"SELECT id, email, password FROM users WHERE email=$1",
		email,
	).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
