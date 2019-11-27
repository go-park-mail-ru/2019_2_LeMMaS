package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/user/model"
	"io"
)

type Repository interface {
	Update(user model.User) error
	UpdateAvatarPath(id int, avatarPath string) error
	GetAll(id int) ([]model.User, error)
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type FileRepository interface {
	Store(file io.Reader) (location string, err error)
}

type SessionRepository interface {
	GetUserBySession(sessionID string) (int, bool)
}