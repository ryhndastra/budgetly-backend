package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(
	db *pgx.Conn,
) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(
	user *User,
) (*User, error) {

	query := `
		INSERT INTO users (
			email,
			password_hash,
			full_name
		)
		VALUES (
			$1, $2, $3
		)
		RETURNING
			id,
			created_at,
			updated_at
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		user.Email,
		user.PasswordHash,
		user.FullName,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetByEmail(
	email string,
) (*User, error) {

	query := `
		SELECT
			id,
			email,
			password_hash,
			full_name,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	var user User

	err := r.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}