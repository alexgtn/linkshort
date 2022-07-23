package pkg

import "errors"

var (
	ErrNotFound    = errors.New("error not found")
	ErrNotSingular = errors.New("error not singular")
)
