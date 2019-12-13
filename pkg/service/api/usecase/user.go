package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"io"
	"strings"
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

func (u *userUsecase) UpdateUser(id int, password, name string) error {
	return nil
}

func (u *userUsecase) UpdateUserAvatar(user *model.User, avatarFile io.Reader) error {
	return nil
}

func (u *userUsecase) GetAvatarUrlByName(name string) string {
	avatarsByName := map[string]string{
		"eath":    "http://www.i2clipart.com/cliparts/3/d/1/e/clipart-earth-3d1e.png",
		"trump":   "https://lemmas.s3.eu-west-3.amazonaws.com/trump.png",
		"lebedev": "https://lemmas.s3.eu-west-3.amazonaws.com/lebedev.jpg",
		"cat":     "https://i.pinimg.com/originals/90/a8/56/90a856d434dd9df24d8d5fdf4bf3ce72.png",
	}
	return avatarsByName[strings.ToLower(name)]
}
