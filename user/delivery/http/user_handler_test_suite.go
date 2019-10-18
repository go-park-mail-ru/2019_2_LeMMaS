package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/test"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

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
	s.handler = NewUserHandler(s.E, s.usecase)
}

func (s *UserHandlerTestSuite) ExpectUsecase() *user.MockUsecaseMockRecorder {
	return s.usecase.EXPECT()
}

func (s *UserHandlerTestSuite) TestUserList(expectedResponse string) {
	s.SetupRequest(http.MethodGet, ApiV1UserListPath, "")
	s.handler.HandleUserList(s.E.NewContext(s.Request, s.Response))
	s.TestResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserRegister(requestBody, expectedResponse string) {
	s.SetupRequest(http.MethodPost, ApiV1UserRegisterPath, requestBody)
	s.handler.HandleUserRegister(s.E.NewContext(s.Request, s.Response))
	s.TestResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserUpdate(requestBody, expectedResponse string) {
	s.SetupRequest(http.MethodPost, ApiV1UserUpdatePath, requestBody)
	s.handler.HandleUserUpdate(s.E.NewContext(s.Request, s.Response))
	s.TestResponse(expectedResponse)
}

func (s *UserHandlerTestSuite) TestUserLogin(requestBody, expectedResponse string, mustHaveSessionCookie bool) {
	s.SetupRequest(http.MethodPost, ApiV1UserLoginPath, requestBody)
	s.handler.HandleUserLogin(s.E.NewContext(s.Request, s.Response))
	s.TestResponse(expectedResponse)
	if mustHaveSessionCookie {
		s.TestCookiePresent(SessionIDCookieName)
	}
}

func (s *UserHandlerTestSuite) TestUserLogout(expectedResponse string) {
	s.SetupRequest(http.MethodPost, ApiV1UserLogoutPath, "")
	s.handler.HandleUserLogout(s.E.NewContext(s.Request, s.Response))
	s.TestResponse(expectedResponse)
	s.TestCookieNotPresent(SessionIDCookieName)
}
