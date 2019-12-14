package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"io"
)

type userUsecase struct {
}

func NewUserUsecase() api.UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	return nil, nil
}

func (u *userUsecase) GetUserByID(userID int) (*model.User, error) {
	return nil, nil
}

func (u *userUsecase) UpdateUser(userID int, password, name string) error {
	return nil
}

func (u *userUsecase) UpdateUserAvatar(userID int, avatarFile io.Reader) error {
	return nil
}

func (u *userUsecase) GetSpecialAvatar(name string) string {
	return ""
}
