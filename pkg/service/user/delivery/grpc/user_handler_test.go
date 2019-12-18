package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c context.Context

func TestUserHandler_Create(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	usecase.EXPECT().Create(test.Email, test.PasswordHash, test.Name).Return(nil)
	res, err := h.Create(c, &user.CreateParams{Email: test.Email, PasswordHash: test.PasswordHash, Name: test.Name})
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}

func TestUserHandler_GetAll(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	users := []*model.User{{ID: test.UserID}}
	usecase.EXPECT().GetAll().Return(users, nil)
	res, err := h.GetAll(c, nil)
	assert.Nil(t, err)
	assert.Equal(t, test.UserID, int(res.Users[0].Id))
}

func TestUserHandler_GetByID(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	u := &model.User{ID: test.UserID}
	usecase.EXPECT().GetByID(test.UserID).Return(u, nil)
	res, err := h.GetByID(c, &user.GetByIDParams{Id: int32(test.UserID)})
	assert.Nil(t, err)
	assert.Equal(t, test.UserID, int(res.User.Id))
}

func TestUserHandler_GetByEmail(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	u := &model.User{Email: test.Email}
	usecase.EXPECT().GetByEmail(test.Email).Return(u, nil)
	res, err := h.GetByEmail(c, &user.GetByEmailParams{Email: test.Email})
	assert.Nil(t, err)
	assert.Equal(t, test.Email, res.User.Email)
}

func TestUserHandler_GetSpecialAvatar(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	avatar := "avatar.jpg"
	usecase.EXPECT().GetSpecialAvatar(test.Name).Return(avatar)
	res, err := h.GetSpecialAvatar(c, &user.GetSpecialAvatarParams{Name: test.Name})
	assert.Nil(t, err)
	assert.Equal(t, avatar, res.AvatarUrl)
}

func TestUserHandler_Update(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	usecase.EXPECT().Update(test.UserID, "", test.Name).Return(nil)
	res, err := h.Update(c, &user.UpdateParams{Id: int32(test.UserID), Name: test.Name})
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}

func TestUserHandler_UpdateAvatar(t *testing.T) {
	usecase := user.NewMockUserUsecase(gomock.NewController(t))
	h := NewUserHandler(usecase)

	avatar := "avatar.jpg"
	usecase.EXPECT().UpdateAvatar(test.UserID, avatar).Return(nil)
	res, err := h.UpdateAvatar(c, &user.UpdateAvatarParams{Id: int32(test.UserID), AvatarPath: avatar})
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}
