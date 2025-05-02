package utils

import "errors"

type ErrorResponse error

var (
	ErrNotFound ErrorResponse = errors.New("not found")
)
