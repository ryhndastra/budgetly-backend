package auth

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Register(
	req RegisterRequest,
) (*User, error) {

	if strings.TrimSpace(req.FullName) == "" {
		return nil, errors.New("full name is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return nil, errors.New("email is required")
	}

	if len(req.Password) < 8 {
		return nil, errors.New(
			"password must be at least 8 characters",
		)
	}

	existingUser, _ := s.repo.GetByEmail(
		req.Email,
	)

	if existingUser != nil {
		return nil, errors.New(
			"email already registered",
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
	}

	return s.repo.Create(user)
}

func (s *Service) Login(
	req LoginRequest,
) (*User, error) {

	user, err := s.repo.GetByEmail(
		req.Email,
	)

	if err != nil {
		return nil, errors.New(
			"invalid email or password",
		)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New(
			"invalid email or password",
		)
	}

	return user, nil
}