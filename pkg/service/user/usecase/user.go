package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/consts"
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

func (u *userUsecase) Create(email string, passwordHash string, name string) error {
	return u.repo.Create(email, passwordHash, name)
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

func (u *userUsecase) Update(id int, passwordHash, name string) error {
	usr, err := u.repo.GetByID(id)
	if err == consts.ErrNotFound {
		return err
	}
	if err != nil {
		return consts.ErrStorageError
	}
	if passwordHash != "" {
		usr.PasswordHash = passwordHash
	}
	if name != "" && usr.Name != name {
		usr.Name = name
		if avatar := u.GetSpecialAvatar(name); avatar != "" {
			usr.AvatarPath = avatar
		}
	}
	return u.repo.Update(usr)
}

func (u *userUsecase) UpdateAvatar(userID int, avatarPath string) error {
	return u.repo.UpdateAvatar(userID, avatarPath)
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
