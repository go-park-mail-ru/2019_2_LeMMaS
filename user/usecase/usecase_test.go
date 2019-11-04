package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

const Password = "ssc-tuatara"
const PasswordHash = "qZxR0nf7CPVg82xpehicuDeDEAH+0Y4uQeWZr8yGZXr1MaRpGtJ+rA"

func TestUserUsecase_GetAllUsers(t *testing.T) {
	userRepo := user.NewMockUserRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil)

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
	userRepo := user.NewMockUserRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil)

	email := "t@mail.ru"
	password := Password
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
	userRepo := user.NewMockUserRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil)

	userToLogin := model.User{Email: "t@mail.ru", PasswordHash: PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	sessionID, err := usecase.Login(userToLogin.Email, Password)
	assert.NotEmpty(t, sessionID)
	assert.Nil(t, err)

	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(nil, nil)
	sessionID, err = usecase.Login(userToLogin.Email, Password)
	assert.Empty(t, sessionID)
	assert.EqualError(t, err, "incorrect email")
}

func TestUserUsecase_GetUserBySessionID(t *testing.T) {
	userRepo := user.NewMockUserRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil)

	userToLogin := model.User{Email: "t@mail.ru", PasswordHash: PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	sessionID, _ := usecase.Login(userToLogin.Email, Password)

	userRepo.EXPECT().GetByID(userToLogin.ID).Return(&userToLogin, nil)
	userBySession, err := usecase.GetUserBySessionID(sessionID)
	assert.NotNil(t, userBySession)
	if userBySession != nil {
		assert.Equal(t, *userBySession, userToLogin)
	}
	assert.Nil(t, err)
}

func TestUserUsecase_Logout(t *testing.T) {
	userRepo := user.NewMockUserRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil)

	userToLogin := model.User{ID: 2, Email: "t@mail.ru", PasswordHash: PasswordHash}
	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
	sessionID, _ := usecase.Login(userToLogin.Email, Password)

	err := usecase.Logout(sessionID)
	assert.Nil(t, err)
	userBySession, _ := usecase.GetUserBySessionID(sessionID)
	assert.Nil(t, userBySession)
}

func TestUserUsecase_GetAvatarUrlByName(t *testing.T) {
	usecase := NewUserUsecase(nil, nil)
	url := usecase.GetAvatarUrlByName("trump")
	assert.Regexp(t, `^http`, url)
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	userRepo := user.NewMockUserRepository(gomock.NewController(t))
	usecase := NewUserUsecase(userRepo, nil)

	oldUser := model.User{ID: 4, Name: "Old Name"}
	userRepo.EXPECT().GetByID(oldUser.ID).Return(&oldUser, nil)
	newUser := model.User{ID: 4, Name: "New Name"}
	userRepo.EXPECT().Update(newUser).Return(nil)
	err := usecase.UpdateUser(oldUser.ID, "", newUser.Name)
	assert.Nil(t, err)
}

func TestUserUsecase_UpdateUserAvatar(t *testing.T) {
	mockController := gomock.NewController(t)
	userRepo := user.NewMockUserRepository(mockController)
	userFileRepo := user.NewMockUserFileRepository(mockController)
	usecase := NewUserUsecase(userRepo, userFileRepo)

	userToUpdate := model.User{ID: 2}
	avatarFile := io.LimitedReader{}
	avatarPath := "1.jpg"
	storageAvatarPath := "2ad2.jpg"
	userFileRepo.EXPECT().StoreAvatar(&userToUpdate, &avatarFile, avatarPath).Return(storageAvatarPath, nil)
	userRepo.EXPECT().UpdateAvatarPath(userToUpdate.ID, storageAvatarPath)
	err := usecase.UpdateUserAvatar(&userToUpdate, &avatarFile, avatarPath)
	assert.Nil(t, err)
}

func TestUserUsecase_GetPasswordHash(t *testing.T) {
	usecase := NewUserUsecase(nil, nil)
	passwordHash := usecase.GetPasswordHash(Password)
	assert.NotEmpty(t, passwordHash)
	assert.True(t, usecase.IsPasswordsEqual(Password, passwordHash))
}
