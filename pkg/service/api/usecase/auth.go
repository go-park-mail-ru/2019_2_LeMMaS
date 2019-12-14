package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
)

type authUsecase struct {
	auth auth.AuthClient
	ctx  context.Context
}

func NewAuthUsecase(auth auth.AuthClient) api.AuthUsecase {
	return &authUsecase{
		auth: auth,
		ctx:  context.Background(),
	}
}

func (u *authUsecase) Register(email, password, name string) error {
	params := &auth.RegisterParams{Email: email, Password: password, Name: name}
	res, err := u.auth.Register(u.ctx, params)
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *authUsecase) Login(email, password string) (sessionID string, err error) {
	params := &auth.LoginParams{Email: email, Password: password}
	res, err := u.auth.Login(u.ctx, params)
	if err != nil {
		return "", err
	}
	if res.Error != "" {
		return "", errors.New(res.Error)
	}
	return res.SessionId, nil
}

func (u *authUsecase) Logout(sessionID string) error {
	params := &auth.LogoutParams{SessionId: sessionID}
	res, err := u.auth.Logout(u.ctx, params)
	if err != nil {
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *authUsecase) GetUserBySession(sessionID string) (*model.User, error) {
	return nil, nil
}
