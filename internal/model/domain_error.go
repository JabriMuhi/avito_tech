package model

import "errors"

var ErrForbidden = errors.New("forbidden")
var ErrNotFound = errors.New("not found")
var ErrUserNotFound = errors.New("error user not found")
