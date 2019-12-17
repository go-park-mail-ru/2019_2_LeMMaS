package usecase

//
//import (
//	"fmt"
//	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
//	user2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
//	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"io"
//	"testing"
//)
//
//func TestUserUsecase_GetAllUsers(t *testing.T) {
//	userRepo := user2.NewMockRepository(gomock.NewController(t))
//	server := NewUserUsecase(userRepo, nil, nil, nil)
//
//	expectedUsers := []model.User{
//		{ID: 4, Email: "test4@mail.ru", Name: "Testik 4"},
//		{ID: 5, Email: "test5@mail.ru", Name: "Testik 5"},
//	}
//	userRepo.EXPECT().GetAll().Return(expectedUsers, nil)
//	users, err := server.GetAll()
//	assert.Nil(t, err, "unexpected error")
//	assert.Equal(t, users, expectedUsers)
//
//	userRepo.EXPECT().GetAll().Return(nil, fmt.Errorf("error"))
//	users, err = server.GetAll()
//	assert.EqualError(t, err, "error")
//	assert.Nil(t, users)
//}
//
//func TestUserUsecase_GetUserByID(t *testing.T) {
//	userRepo := user2.NewMockRepository(gomock.NewController(t))
//	server := NewUserUsecase(userRepo, nil, nil, nil)
//
//	expectedUser := &model.User{ID: 4, Email: "test4@mail.ru", Name: "Testik 4"}
//	userRepo.EXPECT().GetByID(expectedUser.ID).Return(expectedUser, nil)
//	users, err := server.GetByID(expectedUser.ID)
//	assert.Nil(t, err, "unexpected error")
//	assert.Equal(t, users, expectedUser)
//
//	userRepo.EXPECT().GetByID(expectedUser.ID).Return(nil, fmt.Errorf("error"))
//	users, err = server.GetByID(expectedUser.ID)
//	assert.EqualError(t, err, "error")
//	assert.Nil(t, users)
//}
//
//func TestUserUsecase_Register(t *testing.T) {
//	userRepo := user2.NewMockRepository(gomock.NewController(t))
//	server := NewUserUsecase(userRepo, nil, nil, nil)
//
//	email := "t@mail.ru"
//	password := test.Password
//	name := "Test"
//
//	userRepo.EXPECT().GetByEmail(email).Return(nil, nil)
//	userRepo.EXPECT().Create(email, gomock.Any(), name)
//	err := server.Register(email, password, name)
//	assert.Nil(t, err)
//
//	userWithSameEmail := model.User{}
//	userRepo.EXPECT().GetByEmail(email).Return(&userWithSameEmail, nil)
//	err = server.Register(email, password, name)
//	assert.EqualError(t, err, "user with this email already registered")
//}
//
//func TestUserUsecase_Login(t *testing.T) {
//	mockController := gomock.NewController(t)
//	userRepo := user2.NewMockRepository(mockController)
//	sessionRepo := user2.NewMockSessionRepository(mockController)
//	server := NewUserUsecase(userRepo, nil, sessionRepo, nil)
//
//	userToLogin := model.User{ID: 2, Email: "t@mail.ru", PasswordHash: test.PasswordHash}
//	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
//	sessionRepo.EXPECT().Add(gomock.Any(), userToLogin.ID).Return(nil)
//
//	session, err := server.Login(userToLogin.Email, test.Password)
//	assert.NotEmpty(t, session)
//	assert.Nil(t, err)
//
//	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(nil, nil)
//	session, err = server.Login(userToLogin.Email, test.Password)
//	assert.Empty(t, session)
//	assert.EqualError(t, err, "incorrect email")
//}
//
//func TestUserUsecase_GetUserBySessionID(t *testing.T) {
//	mockController := gomock.NewController(t)
//	userRepo := user2.NewMockRepository(mockController)
//	sessionRepo := user2.NewMockSessionRepository(mockController)
//	server := NewUserUsecase(userRepo, nil, sessionRepo, nil)
//
//	userToLogin := model.User{Email: "t@mail.ru", PasswordHash: test.PasswordHash}
//	userRepo.EXPECT().GetByEmail(userToLogin.Email).Return(&userToLogin, nil)
//	userRepo.EXPECT().GetByID(userToLogin.ID).Return(&userToLogin, nil)
//	sessionRepo.EXPECT().Add(gomock.Any(), userToLogin.ID).Return(nil)
//	sessionRepo.EXPECT().Get(gomock.Any()).Return(userToLogin.ID, true)
//
//	session, _ := server.Login(userToLogin.Email, test.Password)
//	userBySession, err := server.GetUserBySession(session)
//	assert.NotNil(t, userBySession)
//	if userBySession != nil {
//		assert.Equal(t, *userBySession, userToLogin)
//	}
//	assert.Nil(t, err)
//}
//
//func TestUserUsecase_Logout(t *testing.T) {
//	sessionRepo := user2.NewMockSessionRepository(gomock.NewController(t))
//	server := NewUserUsecase(nil, nil, sessionRepo, nil)
//	sessionRepo.EXPECT().Delete(test.SessionID).Return(nil)
//	err := server.Logout(test.SessionID)
//	assert.Nil(t, err)
//}
//
//func TestUserUsecase_GetAvatarUrlByName(t *testing.T) {
//	server := NewUserUsecase(nil, nil, nil, nil)
//	url := server.GetSpecialAvatar("trump")
//	assert.Regexp(t, `^http`, url)
//}
//
//func TestUserUsecase_UpdateUser(t *testing.T) {
//	userRepo := user2.NewMockRepository(gomock.NewController(t))
//	server := NewUserUsecase(userRepo, nil, nil, nil)
//
//	oldUser := model.User{ID: 4, Name: "Old Name"}
//	userRepo.EXPECT().GetByID(oldUser.ID).Return(&oldUser, nil)
//	newUser := model.User{ID: 4, Name: "New Name"}
//	userRepo.EXPECT().Update(newUser).Return(nil)
//	err := server.Update(oldUser.ID, "", newUser.Name)
//	assert.Nil(t, err)
//}
//
//func TestUserUsecase_UpdateUserAvatar(t *testing.T) {
//	mockController := gomock.NewController(t)
//	userRepo := user2.NewMockRepository(mockController)
//	userFileRepo := user2.NewMockFileRepository(mockController)
//	server := NewUserUsecase(userRepo, userFileRepo, nil, nil)
//
//	userToUpdate := model.User{ID: 2}
//	avatarFile := io.LimitedReader{}
//	avatarFileLocation := "2ad2.jpg"
//	userFileRepo.EXPECT().Store(&avatarFile).Return(avatarFileLocation, nil)
//	userRepo.EXPECT().UpdateAvatar(userToUpdate.ID, avatarFileLocation)
//	err := server.UpdateAvatar(&userToUpdate, &avatarFile)
//	assert.Nil(t, err)
//}
//
//func TestUserUsecase_GetPasswordHash(t *testing.T) {
//	server := userUsecase{}
//	passwordHash := server.getPasswordHash(test.Password)
//	assert.NotEmpty(t, passwordHash)
//	assert.True(t, server.isPasswordsEqual(test.Password, passwordHash))
//}
