package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/model"
	"io"
)

type Usecase interface {
	GetAllUsers() ([]model.User, error)
	GetUserBySessionID(sessionID string) (*model.User, error)
	UpdateUser(id int, password, name string) error
	UpdateUserAvatar(user *model.User, avatarFile io.Reader, avatarPath string) error
	GetAvatarUrlByName(name string) string
	Register(email, password, name string) error
	Login(email, password string) (sessionID string, err error)
	Logout(sessionID string) error
}
