package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

type UserUsecase interface {
	GetAll() ([]model.User, error)
	GetByID(userID int) (*model.User, error)
	Update(userID int, password, name string) error
	UpdateAvatar(userID int, avatarPath string) error
	GetSpecialAvatar(name string) string
}
