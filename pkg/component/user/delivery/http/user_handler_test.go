package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

var s = NewUserHandlerTestSuite()

func TestUserHandler_HandleUserList(t *testing.T) {
	s.StartTest(t)

	s.ExpectUsecase().GetAllUsers().Return([]model.User{{Name: "Ivan"}}, nil)
	s.TestUserList(`{"status":"ok","body":{"users":[{"id":0,"email":"","name":"Ivan","avatar_path":""}]}}`)
}

func TestUserHandler_HandleUserRegister(t *testing.T) {
	s.StartTest(t)

	user1 := model.User{ID: 1, Email: "testik1@mail.ru", Name: "Test The Best 1"}
	s.ExpectUsecase().Register(user1.Email, test.Password, user1.Name).Return(nil)
	s.ExpectUsecase().GetAllUsers().Return([]model.User{user1}, nil)

	s.TestUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		s.Ok(),
		http.StatusOK,
	)
	s.TestUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1","avatar_path":""}]}}`)

	s.TestUserRegister(
		`{"email": "invalid-email","name": "Test The Best 1","password": "ssc-tuatara"}`,
		s.Error("email: invalid-email does not validate as email"),
		http.StatusBadRequest,
	)

	s.ExpectUsecase().Register(user1.Email, test.Password, user1.Name).Return(fmt.Errorf("user already registered"))
	s.TestUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		s.Error("user already registered"),
		http.StatusBadRequest,
	)
}

func TestUserHandler_HandleUserUpdate(t *testing.T) {
	s.StartTest(t)

	s.ExpectUsecase().Login("testik1@mail.ru", "ssc-tuatara").Return(test.SessionID, nil)
	s.ExpectUsecase().GetUserBySessionID(test.SessionID).Return(&model.User{ID: 1}, nil)
	s.ExpectUsecase().UpdateUser(1, "", "New Name").Return(nil)

	s.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		s.Ok(),
		true,
	)
	s.TestUserUpdate(`{"id":1,"name":"New Name"}`, s.Ok())
}

func TestUserHandler_HandleUserLogin(t *testing.T) {
	s.StartTest(t)
	s.ExpectUsecase().Login("testik1@mail.ru", "ssc-tuatara").Return("", nil)
	s.TestUserLogin(
		`{"email":"testik1@mail.ru","password":"ssc-tuatara"}`,
		s.Ok(),
		true,
	)
}

func TestUserHandler_HandleUserLogout(t *testing.T) {
	s.StartTest(t)

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

type UserHandlerTestSuite struct {
	test.HandlerTestSuite
	usecase *user.MockUsecase
	handler *UserHandler
}

func NewUserHandlerTestSuite() *UserHandlerTestSuite {
	return &UserHandlerTestSuite{
		HandlerTestSuite: *test.NewHandlerTestSuite(),
	}
}

func (s *UserHandlerTestSuite) StartTest(t *testing.T) {
	s.HandlerTestSuite.StartTest(t)
	s.usecase = user.NewMockUsecase(gomock.NewController(t))
	logger := testMock.NewMockLogger(t)
	s.handler = NewUserHandler(s.E, s.usecase, logger)
}

func (s *UserHandlerTestSuite) ExpectUsecase() *user.MockUsecaseMockRecorder {
	return s.usecase.EXPECT()
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
		s.TestCookiePresent(delivery.SessionIDCookieName)
	}
}

func (s *UserHandlerTestSuite) TestUserLogout(expectedResponse string) {
	s.SetupRequestWithBody("")
	s.handler.handleUserLogout(s.NewContext())
	s.TestOkResponse(expectedResponse)
	s.TestCookieNotPresent(delivery.SessionIDCookieName)
}
