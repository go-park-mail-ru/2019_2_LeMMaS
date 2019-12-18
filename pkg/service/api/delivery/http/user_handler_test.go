package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

var us = NewUserHandlerTestSuite()

func TestUserHandler_HandleUserList(t *testing.T) {
	us.StartTest(t)

	us.user.EXPECT().GetAll().Return([]*model.User{{Name: "Ivan"}}, nil)
	us.TestUserList(`{"status":"ok","body":{"users":[{"id":0,"email":"","name":"Ivan","avatar_path":""}]}}`)
}

func TestUserHandler_HandleUserRegister(t *testing.T) {
	us.StartTest(t)

	user1 := model.User{ID: 1, Email: "testik1@mail.ru", Name: "Test The Best 1"}
	us.auth.EXPECT().Register(user1.Email, test.Password, user1.Name).Return(nil)
	us.user.EXPECT().GetAll().Return([]*model.User{&user1}, nil)

	us.TestUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		us.Ok(),
		http.StatusOK,
	)
	us.TestUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1","avatar_path":""}]}}`)

	us.auth.EXPECT().Register(user1.Email, test.Password, user1.Name).Return(fmt.Errorf("user already registered"))
	us.TestUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		us.Error("user already registered"),
		http.StatusBadRequest,
	)
}

func TestUserHandler_HandleUserUpdate(t *testing.T) {
	us.StartTest(t)

	us.auth.EXPECT().Login("testik1@mail.ru", "ssc-tuatara").Return(test.SessionID, nil)
	us.auth.EXPECT().GetUserID(test.SessionID).Return(1, nil)
	us.user.EXPECT().Update(1, "", "New Name").Return(nil)

	us.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		us.Ok(),
		true,
	)
	us.TestUserUpdate(`{"id":1,"name":"New Name"}`, us.Ok())
}

func TestUserHandler_HandleUserLogin(t *testing.T) {
	us.StartTest(t)
	us.auth.EXPECT().Login("testik1@mail.ru", "ssc-tuatara").Return("", nil)
	us.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		us.Ok(),
		true,
	)
}

func TestUserHandler_HandleUserLogout(t *testing.T) {
	us.StartTest(t)

	session := "sess"
	us.auth.EXPECT().Login("testik1@mail.ru", "ssc-tuatara").Return(session, nil)
	us.auth.EXPECT().Logout(session).Return(nil)

	us.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		us.Ok(),
		true,
	)
	us.TestUserLogout(us.Ok())
}

type UserHandlerTestSuite struct {
	test.HandlerTestSuite
	user    *api.MockUserUsecase
	auth    *api.MockAuthUsecase
	handler *UserHandler
}

func NewUserHandlerTestSuite() *UserHandlerTestSuite {
	return &UserHandlerTestSuite{
		HandlerTestSuite: *test.NewHandlerTestSuite(),
	}
}

func (s *UserHandlerTestSuite) StartTest(t *testing.T) {
	s.HandlerTestSuite.StartTest(t)
	controller := gomock.NewController(t)
	s.user = api.NewMockUserUsecase(controller)
	s.auth = api.NewMockAuthUsecase(controller)
	log := testMock.NewMockLogger(t)
	s.handler = NewUserHandler(s.E, s.user, s.auth, log)
}

func (s *UserHandlerTestSuite) TestUserList(expectedResponse string) {
	s.SetupRequestWithBody("")
	s.handler.handleUserList(s.NewContext())
	s.TestOkResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserRegister(requestBody, expectedResponse string, expectedCode int) {
	s.SetupRequestWithBody(requestBody)
	s.handler.handleUserRegister(s.NewContext())
	s.TestResponse(expectedResponse, expectedCode)
}

func (s *UserHandlerTestSuite) TestUserUpdate(requestBody, expectedResponse string) {
	s.SetupRequestWithBody(requestBody)
	s.handler.handleUserUpdate(s.NewContext())
	s.TestOkResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserLogin(requestBody, expectedResponse string, mustHaveSessionCookie bool) {
	s.SetupRequestWithBody(requestBody)
	s.handler.handleUserLogin(s.NewContext())
	s.TestOkResponse(expectedResponse)
	if mustHaveSessionCookie {
		s.TestCookiePresent(delivery.SessionCookieName)
	}
}

func (s *UserHandlerTestSuite) TestUserLogout(expectedResponse string) {
	s.SetupRequestWithBody("")
	s.handler.handleUserLogout(s.NewContext())
	s.TestOkResponse(expectedResponse)
	s.TestCookieNotPresent(delivery.SessionCookieName)
}
