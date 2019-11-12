package http

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

var s = NewUserHandlerTestSuite()

func TestUserHandler_HandleUserList(t *testing.T) {
	s.SetTesting(t)

	s.ExpectUsecase().GetAllUsers().Return([]model.User{{Name: "Ivan"}}, nil)
	s.TestUserList(`{"status":"ok","body":{"users":[{"id":0,"email":"","name":"Ivan","avatar_path":""}]}}`)
}

func TestUserHandler_HandleUserRegister(t *testing.T) {
	s.SetTesting(t)

	user1 := model.User{ID: 1, Email: "testik1@mail.ru", Name: "Test The Best 1"}
	s.ExpectUsecase().Register(user1.Email, test.Password, user1.Name).Return(nil)
	s.ExpectUsecase().GetAllUsers().Return([]model.User{user1}, nil)

	s.TestUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		s.Ok(),
		http.StatusOK,
	)
	s.TestUserList(`{"status":"ok","body":{"users":[{"id":1,"email":"testik1@mail.ru","name":"Test The Best 1","avatar_path":""}]}}`)

	s.ExpectUsecase().Register(user1.Email, test.Password, user1.Name).Return(fmt.Errorf("user already registered"))
	s.TestUserRegister(
		`{"email": "testik1@mail.ru","name": "Test The Best 1","password": "ssc-tuatara"}`,
		s.Error("user already registered"),
		http.StatusBadRequest,
	)
}

func TestUserHandler_HandleUserUpdate(t *testing.T) {
	s.SetTesting(t)

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

func (s *UserHandlerTestSuite) SetTesting(t *testing.T) {
	s.HandlerTestSuite.SetTesting(t)
	s.usecase = user.NewMockUsecase(gomock.NewController(t))
	logger := testMock.NewMockLogger(t)
	s.handler = NewUserHandler(s.E, s.usecase, logger)
}

func (s *UserHandlerTestSuite) ExpectUsecase() *user.MockUsecaseMockRecorder {
	return s.usecase.EXPECT()
}

func (s *UserHandlerTestSuite) TestUserList(expectedResponse string) {
	s.SetupRequest(http.MethodGet, httpDelivery.ApiV1UserListPath, "")
	s.handler.HandleUserList(s.NewHandlerContext())
	s.TestOkResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserRegister(requestBody, expectedResponse string, expectedCode int) {
	s.SetupRequest(http.MethodPost, httpDelivery.ApiV1UserRegisterPath, requestBody)
	s.handler.HandleUserRegister(s.NewHandlerContext())
	s.TestResponse(expectedResponse, expectedCode)
}

func (s *UserHandlerTestSuite) TestUserUpdate(requestBody, expectedResponse string) {
	s.SetupRequest(http.MethodPost, httpDelivery.ApiV1UserUpdatePath, requestBody)
	s.handler.HandleUserUpdate(s.NewHandlerContext())
	s.TestOkResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserLogin(requestBody, expectedResponse string, mustHaveSessionCookie bool) {
	s.SetupRequest(http.MethodPost, httpDelivery.ApiV1UserLoginPath, requestBody)
	s.handler.HandleUserLogin(s.NewHandlerContext())
	s.TestOkResponse(expectedResponse)
	if mustHaveSessionCookie {
		s.TestCookiePresent(httpDelivery.SessionIDCookieName)
	}
}

func (s *UserHandlerTestSuite) TestUserLogout(expectedResponse string) {
	s.SetupRequest(http.MethodPost, httpDelivery.ApiV1UserLogoutPath, "")
	s.handler.HandleUserLogout(s.NewHandlerContext())
	s.TestOkResponse(expectedResponse)
	s.TestCookieNotPresent(httpDelivery.SessionIDCookieName)
}