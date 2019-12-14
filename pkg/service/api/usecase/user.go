package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"io"
)

type userUsecase struct {
	fileRepo api.FileRepository
	user     user.UserClient
	ctx      context.Context
}

func NewUserUsecase(user user.UserClient, fileRepo api.FileRepository) api.UserUsecase {
	return &userUsecase{
		user:     user,
		fileRepo: fileRepo,
		ctx:      context.Background(),
	}
}

func (u *userUsecase) GetAll() ([]*model.User, error) {
	params := user.GetAllParams{}
	res, err := u.user.GetAll(u.ctx, &params)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return u.convertUsers(res.Users), nil
}

func (u *userUsecase) GetByID(id int) (*model.User, error) {
	params := user.GetByIDParams{Id: int32(id)}
	res, err := u.user.GetByID(u.ctx, &params)
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return u.convertUser(res.User), nil
}

func (u *userUsecase) Update(userID int, password, name string) error {
	return nil
}

func (u *userUsecase) UpdateAvatar(id int, avatar io.Reader) error {
	path, err := u.fileRepo.Store(avatar)
	if err != nil {
		return err
	}
	params := user.UpdateAvatarParams{
		UserId:     int32(id),
		AvatarPath: path,
	}
	res, err := u.user.UpdateAvatar(u.ctx, &params)
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *userUsecase) GetSpecialAvatar(name string) string {
	return ""
}

func (u *userUsecase) convertUsers(users []*user.UserData) []*model.User {
	result := make([]*model.User, 0)
	for _, usr := range users {
		result = append(result, u.convertUser(usr))
	}
	return result
}

func (u *userUsecase) convertUser(usr *user.UserData) *model.User {
	return &model.User{
		ID:         int(usr.Id),
		Email:      usr.Email,
		Name:       usr.Name,
		AvatarPath: usr.AvatarPath,
	}
}
