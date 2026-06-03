package category

import "github.com/google/uuid"

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

func (s *Service) GetByUserID(
	userID uuid.UUID,
) ([]Category, error) {
	return s.repo.GetByUserID(userID)
}