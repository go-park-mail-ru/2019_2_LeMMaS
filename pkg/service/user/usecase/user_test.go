package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUsecase_Create(t *testing.T) {
	repo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(repo)

	email := "test1@mail.ru"
	name := "Test"
	repo.EXPECT().Create(email, test.PasswordHash, name).Return(nil)
	err := usecase.Create(email, test.PasswordHash, name)
	assert.Nil(t, err)
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	repo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(repo)

	expectedUsers := []*model.User{
		{ID: 4, Email: "test4@mail.ru", Name: "Testik 4"},
		{ID: 5, Email: "test5@mail.ru", Name: "Testik 5"},
	}
	repo.EXPECT().GetAll().Return(expectedUsers, nil)
	users, err := usecase.GetAll()
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, users, expectedUsers)

	repo.EXPECT().GetAll().Return(nil, fmt.Errorf("error"))
	users, err = usecase.GetAll()
	assert.EqualError(t, err, "error")
	assert.Nil(t, users)
}

func TestUserUsecase_GetUserByID(t *testing.T) {
	repo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(repo)

	expectedUser := &model.User{ID: 4, Email: "test4@mail.ru", Name: "Testik 4"}
	repo.EXPECT().GetByID(expectedUser.ID).Return(expectedUser, nil)
	users, err := usecase.GetByID(expectedUser.ID)
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, users, expectedUser)

	repo.EXPECT().GetByID(expectedUser.ID).Return(nil, fmt.Errorf("error"))
	users, err = usecase.GetByID(expectedUser.ID)
	assert.EqualError(t, err, "error")
	assert.Nil(t, users)
}

func TestUserUsecase_GetAvatarUrlByName(t *testing.T) {
	usecase := NewUserUsecase(nil)
	url := usecase.GetSpecialAvatar("trump")
	assert.Regexp(t, `^http`, url)
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	repo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(repo)

	oldUser := model.User{ID: 4, Name: "Old Name"}
	repo.EXPECT().GetByID(oldUser.ID).Return(&oldUser, nil)
	newUser := &model.User{ID: 4, Name: "New Name"}
	repo.EXPECT().Update(newUser).Return(nil)
	err := usecase.Update(oldUser.ID, "", newUser.Name)
	assert.Nil(t, err)
}

func TestUserUsecase_UpdateUserAvatar(t *testing.T) {
	repo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(repo)

	userID := 2
	path := "2ad2.jpg"
	repo.EXPECT().UpdateAvatar(userID, path)
	err := usecase.UpdateAvatar(userID, path)
	assert.Nil(t, err)
}
