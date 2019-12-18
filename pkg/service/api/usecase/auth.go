package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
)

type authUsecase struct {
	auth auth.AuthClient
	log  logger.Logger

	c context.Context
}

func NewAuthUsecase(auth auth.AuthClient, log logger.Logger) api.AuthUsecase {
	return &authUsecase{
		auth: auth,
		log:  log,
		c:    context.Background(),
	}
}

func (u *authUsecase) Register(email, password, name string) error {
	params := &auth.RegisterParams{Email: email, Password: password, Name: name}
	res, err := u.auth.Register(u.c, params)
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *authUsecase) Login(email, password string) (session string, err error) {
	params := &auth.LoginParams{Email: email, Password: password}
	res, err := u.auth.Login(u.c, params)
	if err != nil {
		u.log.Error(err)
		return "", err
	}
	if res.Error != "" {
		return "", errors.New(res.Error)
	}
	return res.Session, nil
}

func (u *authUsecase) Logout(session string) error {
	params := &auth.LogoutParams{Session: session}
	res, err := u.auth.Logout(u.c, params)
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *authUsecase) GetUserID(session string) (int, error) {
	params := &auth.GetUserParams{Session: session}
	res, err := u.auth.GetUser(u.c, params)
	if err != nil {
		u.log.Error(err)
		return 0, err
	}
	if res.Error != "" {
		return 0, errors.New(res.Error)
	}
	return int(res.Id), nil
}

func (u *authUsecase) GetPasswordHash(password string) (string, error) {
	params := &auth.GetPasswordHashParams{Password: password}
	res, err := u.auth.GetPasswordHash(u.c, params)
	if err != nil {
		u.log.Error(err)
		return "", err
	}
	return res.PasswordHash, nil
}
