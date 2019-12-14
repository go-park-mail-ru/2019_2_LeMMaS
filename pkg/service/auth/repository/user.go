package repository

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
)

type userRepository struct {
	u      user.UserClient
	logger logger.Logger
	ctx    context.Context
}

func NewUserRepository(u user.UserClient, logger logger.Logger) auth.UserRepository {
	return &userRepository{
		u:      u,
		logger: logger,
		ctx:    context.Background(),
	}
}

func (r *userRepository) Register(email string, passwordHash string, name string) error {
	return nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	params := user.GetByEmailParams{Email: email}
	res, err := r.u.GetByEmail(r.ctx, &params)
	if err != nil {
		return nil, err
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
