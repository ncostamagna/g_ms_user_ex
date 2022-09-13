package user

import (
	"log"

	"github.com/ncostamagna/g_ms_user_ex/internal/domain"
)

type (
	Filters struct {
		FirstName string
		LastName  string
	}

	Service interface {
		Create(firstName, lastName, email, phone string) (*domain.User, error)
		Get(id string) (*domain.User, error)
		GetAll(filters Filters, offset, limit int) ([]domain.User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(filters Filters) (int, error)
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
func (s service) Create(firstName, lastName, email, phone string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(user); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return user, nil
}

func (s service) GetAll(filters Filters, offset, limit int) ([]domain.User, error) {

	users, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return users, nil
}

func (s service) Get(id string) (*domain.User, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}
	return user, nil
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s service) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	return s.repo.Update(id, firstName, lastName, email, phone)
}

func (s service) Count(filters Filters) (int, error) {
	return s.repo.Count(filters)
}
