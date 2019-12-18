//go:generate mockgen -source=$GOFILE -destination=repository_mock.go -package=$GOPACKAGE

package api

import "io"

type FileRepository interface {
	Store(file io.Reader) (location string, err error)
}
