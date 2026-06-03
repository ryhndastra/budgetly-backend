package transaction

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

func (r *Repository) Create(
	transaction *Transaction,
) (*Transaction, error) {

	query := `
	INSERT INTO transactions (
		user_id,
		category_id,
		title,
		amount,
		note,
		type
	)
	VALUES (
		$1, $2, $3, $4, $5, $6
	)
	RETURNING
		id,
		created_at,
		updated_at
`

	err := r.db.QueryRow(
		context.Background(),
		query,
		transaction.UserID,
		transaction.CategoryID,
		transaction.Title,
		transaction.Amount,
		transaction.Note,
		transaction.Type,
	).Scan(
		&transaction.ID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *Repository) GetByUserID(
	userID uuid.UUID,
) ([]Transaction, error) {

	query := `
		SELECT
			id,
			user_id,
			category_id,
			title,
			amount,
			note,
			type,
			created_at,
			updated_at
		FROM transactions
		WHERE user_id = $1
		ORDER BY created_at DESC
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

	var transactions []Transaction

	for rows.Next() {
		var transaction Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.CategoryID,
			&transaction.Title,
			&transaction.Amount,
			&transaction.Note,
			&transaction.Type,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(
			transactions,
			transaction,
		)
	}

	return transactions, nil
}
