package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
)

type authUsecase struct {
	auth auth.AuthClient
}

func NewAuthUsecase(auth auth.AuthClient) api.AuthUsecase {
	return &authUsecase{
		auth: auth,
	}
}

func (u *authUsecase) Register(email, password, name string) error {
	userData := &auth.UserDataRegister{email, password, name}
	_, err := u.auth.RegisterUser(context.Background(), userData)
	return err
}

func (u *authUsecase) Login(email, password string) (sessionID string, err error) {
	userData := &auth.UserAuth{email, password}
	result, err := u.auth.Login(context.Background(), userData)
	if err != nil {
		return "", err
	}
	return result.SessionID.ID, err
}

func (u *authUsecase) Logout(sessionID string) error {
	userData := &auth.SessionID{sessionID}
	_, err := u.auth.Logout(context.Background(), userData)
	return err
}

func (u *authUsecase) GetUserBySessionID(sessionID string) (*model.User, error) {
	return nil, nil
}
