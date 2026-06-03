package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByUserID(
	userID uuid.UUID,
) ([]Category, error) {

	query := `
		SELECT
			id,
			user_id,
			name,
			icon,
			color,
			type,
			created_at
		FROM categories
		WHERE user_id = $1
		ORDER BY name ASC
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category

		err := rows.Scan(
				&category.ID,
				&category.UserID,
				&category.Name,
				&category.Icon,
				&category.Color,
				&category.Type,
				&category.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		categories = append(
			categories,
			category,
		)
	}

	return categories, nil
}