package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ac = gomock.Any()

func TestAuthUsecase_GetPasswordHash(t *testing.T) {
	client := auth.NewMockAuthClient(gomock.NewController(t))
	usecase := NewAuthUsecase(client, mock.NewMockLogger(t))

	client.EXPECT().
		GetPasswordHash(ac, &auth.GetPasswordHashParams{Password: test.Password}).
		Return(&auth.GetPasswordHashResult{PasswordHash: test.PasswordHash}, nil)
	passwordHash, err := usecase.GetPasswordHash(test.Password)
	assert.Nil(t, err)
	assert.Equal(t, test.PasswordHash, passwordHash)
}

func TestAuthUsecase_GetUserID(t *testing.T) {
	client := auth.NewMockAuthClient(gomock.NewController(t))
	usecase := NewAuthUsecase(client, mock.NewMockLogger(t))

	client.EXPECT().
		GetUser(ac, &auth.GetUserParams{Session: test.Session}).
		Return(&auth.GetUserResult{Id: test.UserID}, nil)
	id, err := usecase.GetUserID(test.Session)
	assert.Nil(t, err)
	assert.Equal(t, test.UserID, id)
}

func TestAuthUsecase_Login(t *testing.T) {
	client := auth.NewMockAuthClient(gomock.NewController(t))
	usecase := NewAuthUsecase(client, mock.NewMockLogger(t))

	client.EXPECT().
		Login(ac, &auth.LoginParams{Email: test.Email, Password: test.Password}).
		Return(&auth.LoginResult{Session: test.Session}, nil)
	session, err := usecase.Login(test.Email, test.Password)
	assert.Nil(t, err)
	assert.Equal(t, test.Session, session)
}

func TestAuthUsecase_Logout(t *testing.T) {
	client := auth.NewMockAuthClient(gomock.NewController(t))
	usecase := NewAuthUsecase(client, mock.NewMockLogger(t))

	client.EXPECT().
		Logout(ac, &auth.LogoutParams{Session: test.Session}).
		Return(&auth.LogoutResult{Error: ""}, nil)
	err := usecase.Logout(test.Session)
	assert.Nil(t, err)
}

func TestAuthUsecase_Register(t *testing.T) {
	client := auth.NewMockAuthClient(gomock.NewController(t))
	usecase := NewAuthUsecase(client, mock.NewMockLogger(t))

	client.EXPECT().
		Register(ac, &auth.RegisterParams{Email: test.Email, Password: test.Password, Name: test.Name}).
		Return(&auth.RegisterResult{Error: ""}, nil)
	err := usecase.Register(test.Email, test.Password, test.Name)
	assert.Nil(t, err)
}

func TestAuthUsecase_(t *testing.T) {
	client := auth.NewMockAuthClient(gomock.NewController(t))
	usecase := NewAuthUsecase(client, mock.NewMockLogger(t))

	client.EXPECT().
		Register(ac, &auth.RegisterParams{Email: test.Email, Password: test.Password, Name: test.Name}).
		Return(&auth.RegisterResult{Error: ""}, nil)
	err := usecase.Register(test.Email, test.Password, test.Name)
	assert.Nil(t, err)
}
