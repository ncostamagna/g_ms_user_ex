package user

import (
	"errors"
	"fmt"
)

var ErrFirstNameRequired = errors.New("first name is required")
var ErrLastNameRequired = errors.New("last name is required")

// vamos a ir y ver el objeto error de nuevo
type ErrNotFound struct {
	UserID string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("user '%s' doesn't exist", e.UserID)
}
