package transaction

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(
	repo *Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	transaction *Transaction,
) (*Transaction, error) {

	if strings.TrimSpace(transaction.Title) == "" {
		return nil, errors.New("title is required")
	}

	if transaction.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	if transaction.Type != "income" &&
		transaction.Type != "expense" {
		return nil, errors.New("invalid transaction type")
	}

	return s.repo.Create(transaction)
}

func (s *Service) GetByUserID(
	userID uuid.UUID,
) ([]Transaction, error) {
	return s.repo.GetByUserID(userID)
}