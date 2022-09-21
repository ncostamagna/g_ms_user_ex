package user

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ncostamagna/g_ms_user_ex/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error)
		Get(ctx context.Context, id string) (*domain.User, error)
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(ctx context.Context, filters Filters) (int, error)
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

func (r *repo) Create(ctx context.Context, user *domain.User) error {

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		r.log.Println(err)
		return err
	}
	r.log.Println("user created with id: ", user.ID)
	return nil
}

func (r *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.User, error) {
	var u []domain.User

	tx := r.db.WithContext(ctx).Model(&u)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&u)

	if result.Error != nil {
		r.log.Println(result.Error)
		return nil, result.Error
	}
	return u, nil
}

func (r *repo) Get(ctx context.Context, id string) (*domain.User, error) {
	user := domain.User{ID: id}
	result := r.db.WithContext(ctx).First(&user)

	if result.Error != nil {
		r.log.Println(result.Error)
		return nil, ErrNotFound
	}
	return &user, nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	user := domain.User{ID: id}
	result := r.db.WithContext(ctx).Delete(&user)

	if result.Error != nil {
		r.log.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Printf("user %s doesn't exists", id)
		return ErrNotFound
	}
	return nil
}

// PATCH vs PUT
func (r *repo) Update(ctx context.Context, id string, firstName *string, lastName *string, email *string, phone *string) error {

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

	result := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		r.log.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.Printf("user %s doesn't exists", id)
		return ErrNotFound
	}

	return nil
}

func (r *repo) Count(ctx context.Context, filters Filters) (int, error) {
	var count int64
	tx := r.db.WithContext(ctx).Model(domain.User{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		r.log.Println(err)
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
