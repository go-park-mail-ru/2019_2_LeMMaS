package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
)

type userRepository struct {
	u   user.UserClient
	log logger.Logger
	c   context.Context
}

func NewUserRepository(u user.UserClient, log logger.Logger) auth.UserRepository {
	return &userRepository{
		u:   u,
		log: log,
		c:   context.Background(),
	}
}

func (r *userRepository) Create(email string, passwordHash string, name string) error {
	params := user.CreateParams{
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
	}
	res, err := r.u.Create(r.c, &params)
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	params := user.GetByEmailParams{Email: email}
	res, err := r.u.GetByEmail(r.c, &params)
	if err != nil {
		return nil, err
	}
	if res.Error == consts.ErrNotFound.Error() {
		return nil, nil
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}
	return r.convertUser(res.User), nil
}

func (r *userRepository) convertUser(usr *user.UserData) *model.User {
	return &model.User{
		ID:           int(usr.Id),
		Email:        usr.Email,
		PasswordHash: usr.PasswordHash,
		Name:         usr.Name,
		AvatarPath:   usr.AvatarPath,
	}
}
