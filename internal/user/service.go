package user

import (
	"fmt"

	"context"

	"github.com/ncostamagna/g_ms_user_ex/internal/domain"
	"github.com/ncostamagna/g_ms_user_ex/pkg/logger"
)

type (
	Filters struct {
		FirstName string
		LastName  string
	}

	Service interface {
		Create(ctx context.Context, firstName, lastName, email, phone string) (*domain.User, error)
		Get(ctx context.Context, id string) (*domain.User, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error)
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(ctx context.Context, filters Filters) (int, error)
	}

	service struct {
		log  logger.Logger
		repo Repository
	}
)

//NewService is a service handler
func NewService(l logger.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

//Create service
func (s service) Create(ctx context.Context, firstName, lastName, email, phone string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		logger.Error(s.log, err.Error())
		return nil, err
	}

	logger.Success(s.log, fmt.Sprintf("user created with id: %s", user.ID))
	return user, nil
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error) {

	users, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		logger.Error(s.log, err.Error())
		return nil, err
	}
	return users, nil
}

func (s service) Get(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		logger.Error(s.log, err.Error())
		return nil, err
	}
	return user, nil
}

func (s service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s service) Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error {
	return s.repo.Update(ctx, id, firstName, lastName, email, phone)
}

func (s service) Count(ctx context.Context, filters Filters) (int, error) {
	return s.repo.Count(ctx, filters)
}
