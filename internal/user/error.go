package user

import "errors"

var ErrNotFound = errors.New("user not found")
var ErrFirstNameRequired = errors.New("first name is required")
var ErrLastNameRequired = errors.New("last name is required")
