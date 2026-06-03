package transaction

type CreateTransactionRequest struct {
	UserID     string  `json:"user_id"`
	CategoryID string  `json:"category_id"`

	Title string  `json:"title"`
	Amount float64 `json:"amount"`
	Note string    `json:"note"`
	Type string    `json:"type"`
}