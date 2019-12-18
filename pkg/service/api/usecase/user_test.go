package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

var uc = gomock.Any()

func TestUserUsecase_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := user.NewMockUserClient(ctrl)
	usecase := NewUserUsecase(client, nil, mock.NewMockLogger(t))

	users := []*model.User{{ID: test.UserID}}
	expected := []*user.UserData{{Id: test.UserID}}
	client.EXPECT().GetAll(uc, gomock.Any()).Return(&user.GetAllResult{Users: expected}, nil)
	res, err := usecase.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, users, res)
}

func TestUserUsecase_GetByID(t *testing.T) {
	controller := gomock.NewController(t)
	client := user.NewMockUserClient(controller)
	usecase := NewUserUsecase(client, nil, mock.NewMockLogger(t))

	usr := model.User{ID: test.UserID}
	client.EXPECT().
		GetByID(uc, &user.GetByIDParams{Id: test.UserID}).
		Return(&user.GetByIDResult{User: &user.UserData{Id: test.UserID}}, nil)
	res, err := usecase.GetByID(test.UserID)
	assert.Nil(t, err)
	assert.Equal(t, usr.ID, res.ID)
}

func TestUserUsecase_GetSpecialAvatar(t *testing.T) {
	controller := gomock.NewController(t)
	client := user.NewMockUserClient(controller)
	usecase := NewUserUsecase(client, nil, mock.NewMockLogger(t))

	avatar := "avatar.jpg"
	client.EXPECT().
		GetSpecialAvatar(uc, &user.GetSpecialAvatarParams{Name: test.Name}).
		Return(&user.GetSpecialAvatarResult{AvatarUrl: avatar}, nil)
	res, err := usecase.GetSpecialAvatar(test.Name)
	assert.Nil(t, err)
	assert.Equal(t, avatar, res)
}

func TestUserUsecase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := user.NewMockUserClient(ctrl)
	usecase := NewUserUsecase(client, nil, mock.NewMockLogger(t))

	client.EXPECT().
		Update(uc, &user.UpdateParams{Id: test.UserID, Name: test.Name}).
		Return(&user.UpdateResult{Error: ""}, nil)
	err := usecase.Update(test.UserID, "", test.Name)
	assert.Nil(t, err)
}

func TestUserUsecase_UpdateAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := user.NewMockUserClient(ctrl)
	fileRepo := api.NewMockFileRepository(ctrl)
	usecase := NewUserUsecase(client, fileRepo, mock.NewMockLogger(t))

	location := "files.com/avatar.jpg"
	file := &io.LimitedReader{}
	client.EXPECT().
		UpdateAvatar(uc, &user.UpdateAvatarParams{Id: test.UserID, AvatarPath: location}).
		Return(&user.UpdateAvatarResult{Error: ""}, nil)
	fileRepo.EXPECT().Store(file).Return(location, nil)
	err := usecase.UpdateAvatar(test.UserID, file)
	assert.Nil(t, err)
}
