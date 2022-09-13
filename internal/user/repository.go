package user

import (
	"fmt"
	"log"
	"strings"

	"github.com/ncostamagna/g_ms_user_ex/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(user *domain.User) error
		GetAll(filters Filters, offset, limit int) ([]domain.User, error)
		Get(id string) (*domain.User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(filters Filters) (int, error)
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

//NewRepo is a repositories handler
func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(user *domain.User) error {

	if err := r.db.Create(user).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("user created with id: ", user.ID)
	return nil
}

func (r *repo) GetAll(filters Filters, offset, limit int) ([]domain.User, error) {
	var u []domain.User

	tx := r.db.Model(&u)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&u)

	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (r *repo) Get(id string) (*domain.User, error) {
	user := domain.User{ID: id}
	result := r.db.First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repo) Delete(id string) error {
	user := domain.User{ID: id}
	result := r.db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// PATCH vs PUT
func (r *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {

	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}
	if phone != nil {
		values["phone"] = *phone
	}

	if err := r.db.Model(&domain.User{}).Where("id = ?", id).Updates(values); err.Error != nil {
		return err.Error
	}

	return nil
}

func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(domain.User{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}
	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}

	return tx
}
