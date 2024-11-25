package repository

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrDuplicate     = errors.New("duplicate")
	ErrNotExists     = errors.New("not exists")
	ErrUnimplemented = errors.New("unimplemented")
)
