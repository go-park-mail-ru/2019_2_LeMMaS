package user

import "github.com/go-park-mail-ru/2019_2_LeMMaS/model"

type Repository interface {
	Create(email string, passwordHash string, name string)
	Update(id int, passwordHash string, name string)
	UpdateAvatarPath(id int, avatarPath string)
	GetAll() []model.User
	GetByID(id int) *model.User
	GetByEmail(email string) *model.User
}
