package user

import (
	"context"
	"log"

	"github.com/ncostamagna/g_ms_user_ex/internal/domain"
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
		log  *log.Logger
		repo Repository
	}
)

//NewService is a service handler
func NewService(l *log.Logger, repo Repository) Service {
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
		s.log.Println(err)
		return nil, err
	}

	return user, nil
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return users, nil
}

func (s service) Get(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return user, nil
}

func (s service) Delete(ctx context.Context, id string) error {

	if err := s.repo.Delete(ctx, id); err != nil {
		s.log.Println(err)
		return err
	}

	return nil
}

// 2 formas de validar si exite, mediante el get o mediante el result del repository
// ver que podemos tener problemas al agregar texto mayor a lo que se espera en la base de datos
func (s service) Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error {

	if err := s.repo.Update(ctx, id, firstName, lastName, email, phone); err != nil {
		s.log.Println(err)
		return err
	}
	return nil
}

func (s service) Count(ctx context.Context, filters Filters) (int, error) {
	return s.repo.Count(ctx, filters)
}
