package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	CategoryID uuid.UUID `json:"category_id"`

	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	Note   string  `json:"note"`
	Type   string  `json:"type"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}