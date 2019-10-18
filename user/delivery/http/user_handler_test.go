package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"testing"
)

var s = NewUserHandlerTestSuite()

func TestUserHandler_HandleUserList(t *testing.T) {
	s.SetTesting(t)

	s.ExpectUsecase().GetAllUsers().Return([]model.User{
		{Name: "Ivan"},
	})
	s.TestUserList(`{"status":"ok","body":{"users":[{"id":0,"email":"","name":"Ivan","avatar_path":""}]}}`)
}

func TestUserHandler_HandleUserRegister(t *testing.T) {
	s.SetTesting(t)

	user1 := model.User{ID: 1, Email: "testik1@mail.ru", Name: "Test The Best 1"}
	password := "ssc-tuatara"
	s.ExpectUsecase().Register(user1.Email, password, user1.Name).Return(nil)
	s.ExpectUsecase().GetAllUsers().Return([]model.User{user1})

	s.TestUserRegister(`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`, s.Ok())
	s.TestUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1","avatar_path":""}]}}`)
}

func TestUserHandler_HandleUserUpdate(t *testing.T) {
	s.SetTesting(t)

	sessionID := "sess"
	s.ExpectUsecase().Login("testik1@mail.ru", "ssc-tuatara").Return(sessionID, nil)
	s.ExpectUsecase().GetUserBySessionID(sessionID).Return(&model.User{ID: 1})
	s.ExpectUsecase().UpdateUser(1, "", "New Name").Return()

	s.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		s.Ok(),
		true,
	)
	s.TestUserUpdate(`{"id":1,"name":"New Name"}`, s.Ok())
}

func TestUserHandler_HandleUserLogin(t *testing.T) {
	s.SetTesting(t)
	s.ExpectUsecase().Login("testik1@mail.ru", "ssc-tuatara").Return("", nil)
	s.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		s.Ok(),
		true,
	)
}

func TestUserHandler_HandleUserLogout(t *testing.T) {
	s.SetTesting(t)

	sessionID := "sess"
	s.ExpectUsecase().Login("testik1@mail.ru", "ssc-tuatara").Return(sessionID, nil)
	s.ExpectUsecase().Logout(sessionID).Return(nil)

	s.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		s.Ok(),
		true,
	)
	s.TestUserLogout(s.Ok())
}
