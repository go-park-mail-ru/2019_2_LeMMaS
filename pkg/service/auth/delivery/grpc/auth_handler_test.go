package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c context.Context = nil

func TestAuthHandler_Login(t *testing.T) {
	usecase := auth.NewMockAuthUsecase(gomock.NewController(t))
	h := NewAuthHandler(usecase)

	usecase.EXPECT().Login(test.Email, test.Password).Return(test.Session, nil)
	res, err := h.Login(c, &auth.LoginParams{
		Email:    test.Email,
		Password: test.Password,
	})
	assert.Equal(t, &auth.LoginResult{Session: test.Session}, res)
	assert.Nil(t, err)
}

func TestAuthHandler_Logout(t *testing.T) {
	usecase := auth.NewMockAuthUsecase(gomock.NewController(t))
	h := NewAuthHandler(usecase)

	usecase.EXPECT().Logout(test.Session).Return(nil)
	res, err := h.Logout(c, &auth.LogoutParams{Session: test.Session})
	assert.Equal(t, &auth.LogoutResult{Error: ""}, res)
	assert.Nil(t, err)
}

func TestAuthHandler_Register(t *testing.T) {
	usecase := auth.NewMockAuthUsecase(gomock.NewController(t))
	h := NewAuthHandler(usecase)

	name := "Artem"
	usecase.EXPECT().Register(test.Email, test.Password, name).Return(nil)
	res, err := h.Register(c, &auth.RegisterParams{
		Email:    test.Email,
		Password: test.Password,
		Name:     name,
	})
	assert.Equal(t, &auth.RegisterResult{Error: ""}, res)
	assert.Nil(t, err)
}

func TestAuthHandler_GetUser(t *testing.T) {
	usecase := auth.NewMockAuthUsecase(gomock.NewController(t))
	h := NewAuthHandler(usecase)

	id := 5
	usecase.EXPECT().GetUser(test.Session).Return(id, true)
	res, err := h.GetUser(c, &auth.GetUserParams{Session: test.Session})
	assert.Equal(t, &auth.GetUserResult{Id: int32(id)}, res)
	assert.Nil(t, err)

	usecase.EXPECT().GetUser(test.Session).Return(0, false)
	res, err = h.GetUser(c, &auth.GetUserParams{Session: test.Session})
	assert.Equal(t, &auth.GetUserResult{Error: consts.ErrNotFound.Error()}, res)
	assert.Nil(t, err)
}

func TestAuthHandler_GetPasswordHash(t *testing.T) {
	usecase := auth.NewMockAuthUsecase(gomock.NewController(t))
	h := NewAuthHandler(usecase)

	usecase.EXPECT().GetPasswordHash(test.Password).Return(test.PasswordHash)
	res, err := h.GetPasswordHash(c, &auth.GetPasswordHashParams{Password: test.Password})
	assert.Equal(t, &auth.GetPasswordHashResult{PasswordHash: test.PasswordHash}, res)
	assert.Nil(t, err)
}
