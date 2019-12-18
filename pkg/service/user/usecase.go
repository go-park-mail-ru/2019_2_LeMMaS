//go:generate mockgen -source=$GOFILE -destination=usecase_mock.go -package=$GOPACKAGE

package user

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

type UserUsecase interface {
	Create(email string, passwordHash string, name string) error
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(id int, passwordHash, name string) error
	UpdateAvatar(id int, avatarPath string) error
	GetSpecialAvatar(name string) string
}
