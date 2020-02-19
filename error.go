package bqworker

import (
	"errors"
)

var (
	ErrInvalidParams = errors.New("[err] invalid params")
	ErrTimeout       = errors.New("[err] timeout")
)
