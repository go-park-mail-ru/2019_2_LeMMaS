//go:generate mockgen -source=$GOFILE -destination=repository_mock.go -package=$GOPACKAGE

package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"io"
)

type Repository interface {
	Create(email string, passwordHash string, name string) error
	Update(user model.User) error
	UpdateAvatarPath(id int, avatarPath string) error
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type FileRepository interface {
	Store(file io.Reader) (location string, err error)
}

type SessionRepository interface {
	AddSession(sessionID string, userID int) error
	GetUserBySession(sessionID string) (int, bool)
	DeleteSession(sessionID string) error
}
