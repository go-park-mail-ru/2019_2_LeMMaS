package consts

import "errors"

var (
	ErrStorageError = errors.New("storage error")
	ErrNotFound     = errors.New("not found")
)
