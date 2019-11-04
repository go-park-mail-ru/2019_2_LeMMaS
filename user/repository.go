package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"io"
)

type UserRepository interface {
	Create(email string, passwordHash string, name string) error
	Update(user model.User) error
	UpdateAvatarPath(id int, avatarPath string) error
	GetAll() ([]model.User, error)
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type UserFileRepository interface {
	StoreAvatar(user *model.User, avatarFile io.Reader, avatarPath string) (string, error)
}
