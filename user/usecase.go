package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"io"
)

type Usecase interface {
	GetAllUsers() []model.User
	GetUserBySessionID(sessionID string) *model.User
	UpdateUser(id int, password, name string)
	UpdateUserAvatar(user *model.User, avatarFile io.Reader, avatarPath string) error
	Register(email, password, name string) error
	Login(email, password string) (string, error)
	Logout(sessionID string) error
}
