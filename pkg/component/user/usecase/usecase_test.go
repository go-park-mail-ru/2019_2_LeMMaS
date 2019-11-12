package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestUserUsecase_GetAllUsers(t *testing.T) {
	userRepo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil, nil)

	expectedUsers := []model.User{
		{ID: 4, Email: "test4@mail.ru", Name: "Testik 4"},
		{ID: 5, Email: "test5@mail.ru", Name: "Testik 5"},
	}
	userRepo.EXPECT().GetAll().Return(expectedUsers, nil)
	users, err := usecase.GetAllUsers()
	assert.Nil(t, err, "unexpected error")
	assert.Equal(t, users, expectedUsers)

	userRepo.EXPECT().GetAll().Return(nil, fmt.Errorf("error"))
	users, err = usecase.GetAllUsers()
	assert.EqualError(t, err, "error")
	assert.Nil(t, users)
}

func TestUserUsecase_Register(t *testing.T) {
	userRepo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil, nil)

	email := "t@mail.ru"
	password := test.Password
	name := "Test"

	userRepo.EXPECT().GetByEmail(email).Return(nil, nil)
	userRepo.EXPECT().Create(email, gomock.Any(), name)
	err := usecase.Register(email, password, name)
	assert.Nil(t, err)

	userWithSameEmail := model.User{}
	userRepo.EXPECT().GetByEmail(email).Return(&userWithSameEmail, nil)
	err = usecase.Register(email, password, name)
	assert.EqualError(t, err, fmt.Sprintf("user with email %v already registered", email))
}

func TestUserUsecase_Login(t *testing.T) {
	mockController := gomock.NewController(t)
	userRepo := user.NewMockRepository(mockController)
	sessionRepo := user.NewMockSessionRepository(mockController)
	usecase := NewUserUsecase(userRepo, nil, sessionRepo)

	userToLogin := model.User{ID: 2, Email: "t@mail.ru", PasswordHash: test.PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	sessionRepo.EXPECT().AddSession(gomock.Any(), userToLogin.ID).Return(nil)

	sessionID, err := usecase.Login(userToLogin.Email, test.Password)
	assert.NotEmpty(t, sessionID)
	assert.Nil(t, err)

	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(nil, nil)
	sessionID, err = usecase.Login(userToLogin.Email, test.Password)
	assert.Empty(t, sessionID)
	assert.EqualError(t, err, "incorrect email")
}

func TestUserUsecase_GetUserBySessionID(t *testing.T) {
	mockController := gomock.NewController(t)
	userRepo := user.NewMockRepository(mockController)
	sessionRepo := user.NewMockSessionRepository(mockController)
	usecase := NewUserUsecase(userRepo, nil, sessionRepo)

	userToLogin := model.User{Email: "t@mail.ru", PasswordHash: test.PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	userRepo.EXPECT().GetByID(userToLogin.ID).Return(&userToLogin, nil)
	sessionRepo.EXPECT().AddSession(gomock.Any(), userToLogin.ID).Return(nil)
	sessionRepo.EXPECT().GetUserBySession(gomock.Any()).Return(userToLogin.ID, true)

	sessionID, _ := usecase.Login(userToLogin.Email, test.Password)
	userBySession, err := usecase.GetUserBySessionID(sessionID)
	assert.NotNil(t, userBySession)
	if userBySession != nil {
		assert.Equal(t, *userBySession, userToLogin)
	}
	assert.Nil(t, err)
}

func TestUserUsecase_Logout(t *testing.T) {
	sessionRepo := user.NewMockSessionRepository(gomock.NewController(t))
	usecase := NewUserUsecase(nil, nil, sessionRepo)
	sessionRepo.EXPECT().DeleteSession(test.SessionID).Return(nil)
	err := usecase.Logout(test.SessionID)
	assert.Nil(t, err)
}

func TestUserUsecase_GetAvatarUrlByName(t *testing.T) {
	usecase := NewUserUsecase(nil, nil, nil)
	url := usecase.GetAvatarUrlByName("trump")
	assert.Regexp(t, `^http`, url)
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	userRepo := user.NewMockRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil, nil)

	oldUser := model.User{ID: 4, Name: "Old Name"}
	userRepo.EXPECT().GetByID(oldUser.ID).Return(&oldUser, nil)
	newUser := model.User{ID: 4, Name: "New Name"}
	userRepo.EXPECT().Update(newUser).Return(nil)
	err := usecase.UpdateUser(oldUser.ID, "", newUser.Name)
	assert.Nil(t, err)
}

func TestUserUsecase_UpdateUserAvatar(t *testing.T) {
	mockController := gomock.NewController(t)
	userRepo := user.NewMockRepository(mockController)
	userFileRepo := user.NewMockFileRepository(mockController)
	usecase := NewUserUsecase(userRepo, userFileRepo, nil)

	userToUpdate := model.User{ID: 2}
	avatarFile := io.LimitedReader{}
	avatarFileLocation := "2ad2.jpg"
	userFileRepo.EXPECT().Store(&avatarFile).Return(avatarFileLocation, nil)
	userRepo.EXPECT().UpdateAvatarPath(userToUpdate.ID, avatarFileLocation)
	err := usecase.UpdateUserAvatar(&userToUpdate, &avatarFile)
	assert.Nil(t, err)
}

func TestUserUsecase_GetPasswordHash(t *testing.T) {
	usecase := userUsecase{}
	passwordHash := usecase.getPasswordHash(test.Password)
	assert.NotEmpty(t, passwordHash)
	assert.True(t, usecase.isPasswordsEqual(test.Password, passwordHash))
}
