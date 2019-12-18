package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthUsecase_GetPasswordHash(t *testing.T) {
	usecase := authUsecase{}
	passwordHash := usecase.GetPasswordHash(test.Password)
	assert.NotEmpty(t, passwordHash)
	assert.True(t, usecase.isPasswordsEqual(test.Password, passwordHash))
}

func TestUserUsecase_Login(t *testing.T) {
	controller := gomock.NewController(t)
	userRepo := auth.NewMockUserRepository(controller)
	sessionRepo := auth.NewMockSessionRepository(controller)
	usecase := NewAuthUsecase(userRepo, sessionRepo, mock.NewMockLogger(t))

	userToLogin := model.User{ID: 2, Email: "t@mail.ru", PasswordHash: test.PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	sessionRepo.EXPECT().Add(gomock.Any(), userToLogin.ID).Return(nil)

	session, err := usecase.Login(userToLogin.Email, test.Password)
	assert.NotEmpty(t, session)
	assert.Nil(t, err)

	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(nil, nil)
	session, err = usecase.Login(userToLogin.Email, test.Password)
	assert.Empty(t, session)
	assert.EqualError(t, err, "incorrect email")
}

func TestUserUsecase_GetUserBySessionID(t *testing.T) {
	controller := gomock.NewController(t)
	userRepo := auth.NewMockUserRepository(controller)
	sessionRepo := auth.NewMockSessionRepository(controller)
	usecase := NewAuthUsecase(userRepo, sessionRepo, mock.NewMockLogger(t))

	userToLogin := model.User{Email: "t@mail.ru", PasswordHash: test.PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	sessionRepo.EXPECT().Add(gomock.Any(), userToLogin.ID).Return(nil)
	sessionRepo.EXPECT().Get(gomock.Any()).Return(userToLogin.ID, true)

	session, _ := usecase.Login(userToLogin.Email, test.Password)
	loggedInUser, ok := usecase.GetUser(session)
	if assert.True(t, ok) {
		assert.Equal(t, loggedInUser, userToLogin.ID)
	}
}

func TestUserUsecase_Logout(t *testing.T) {
	sessionRepo := auth.NewMockSessionRepository(gomock.NewController(t))
	usecase := NewAuthUsecase(nil, sessionRepo, mock.NewMockLogger(t))
	sessionRepo.EXPECT().Delete(test.SessionID).Return(nil)
	err := usecase.Logout(test.SessionID)
	assert.Nil(t, err)
}

func TestUserUsecase_Register(t *testing.T) {
	userRepo := auth.NewMockUserRepository(gomock.NewController(t))
	usecase := NewAuthUsecase(userRepo, nil, mock.NewMockLogger(t))

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
	assert.EqualError(t, err, "user with this email already registered")
}
