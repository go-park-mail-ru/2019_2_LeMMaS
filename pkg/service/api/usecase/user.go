package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"io"
)

type userUsecase struct {
	fileRepo api.FileRepository
	user     user.UserClient
	logger   logger.Logger

	c context.Context
}

func NewUserUsecase(user user.UserClient, fileRepo api.FileRepository, logger logger.Logger) api.UserUsecase {
	return &userUsecase{
		user:     user,
		fileRepo: fileRepo,
		logger:   logger,
		c:        context.Background(),
	}
}

func (u *userUsecase) GetAll() ([]*model.User, error) {
	params := user.GetAllParams{}
	res, err := u.user.GetAll(u.c, &params)
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return u.convertUsers(res.Users), nil
}

func (u *userUsecase) GetByID(id int) (*model.User, error) {
	params := user.GetByIDParams{Id: int32(id)}
	res, err := u.user.GetByID(u.c, &params)
	if err != nil {
		u.logger.Error(err)
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return u.convertUser(res.User), nil
}

func (u *userUsecase) Update(id int, passwordHash, name string) error {
	params := user.UpdateParams{
		Id:           int32(id),
		PasswordHash: passwordHash,
		Name:         name,
	}
	res, err := u.user.Update(u.c, &params)
	if err != nil {
		u.logger.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *userUsecase) UpdateAvatar(id int, avatar io.Reader) error {
	path, err := u.fileRepo.Store(avatar)
	if err != nil {
		return err
	}
	params := user.UpdateAvatarParams{
		Id:         int32(id),
		AvatarPath: path,
	}
	res, err := u.user.UpdateAvatar(u.c, &params)
	if err != nil {
		u.logger.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *userUsecase) GetSpecialAvatar(name string) (string, error) {
	params := user.GetSpecialAvatarParams{Name: name}
	res, err := u.user.GetSpecialAvatar(u.c, &params)
	if err != nil {
		u.logger.Error(err)
		return "", err
	}
	return res.AvatarUrl, nil
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
