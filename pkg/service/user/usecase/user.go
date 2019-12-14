package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"strings"
)

type userUsecase struct {
	repo user.Repository
}

func NewUserUsecase(repo user.Repository) user.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) GetAll() ([]*model.User, error) {
	return u.repo.GetAll()
}

func (u *userUsecase) GetByID(id int) (*model.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) GetByEmail(email string) (*model.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *userUsecase) Update(userID int, password, name string) error {
	return nil
	//userToUpdate, err := u.repository.GetByID(int(pbUserToUpdate.UserID))
	//if err != nil {
	//	return &user.Error{"unknown error"}, err
	//}
	//if pbUserToUpdate.Password != "" {
	//	userToUpdate.PasswordHash = u.getPasswordHash(pbUserToUpdate.Password)
	//}
	//if pbUserToUpdate.Name != "" && userToUpdate.Name != pbUserToUpdate.Name {
	//	userToUpdate.Name = pbUserToUpdate.Name
	//	avatarPath, _ := u.GetAvatarUrlByName(ctx, &user.UserName{pbUserToUpdate.Name})
	//	if avatarPath.AvatarUrl != "" {
	//		userToUpdate.AvatarPath = avatarPath.AvatarUrl
	//	}
	//}
	//err = u.repository.Update(*userToUpdate)
	//if err != nil {
	//	return &user.Error{"unknown error"}, err
	//}
	//return &user.Error{"ok"}, nil
}

func (u *userUsecase) UpdateAvatar(userID int, avatarPath string) error {
	return nil
}

func (u *userUsecase) GetSpecialAvatar(name string) string {
	avatarsByName := map[string]string{
		"earth":   "http://www.i2clipart.com/cliparts/3/d/1/e/clipart-earth-3d1e.png",
		"trump":   "https://lemmas.s3.eu-west-3.amazonaws.com/trump.png",
		"lebedev": "https://lemmas.s3.eu-west-3.amazonaws.com/lebedev.jpg",
		"cat":     "https://i.pinimg.com/originals/90/a8/56/90a856d434dd9df24d8d5fdf4bf3ce72.png",
	}
	return avatarsByName[strings.ToLower(name)]
}
