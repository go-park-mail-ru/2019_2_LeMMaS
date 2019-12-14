package api

import "io"

type FileRepository interface {
	Store(file io.Reader) (location string, err error)
}
