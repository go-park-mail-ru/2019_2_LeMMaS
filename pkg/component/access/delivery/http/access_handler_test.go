package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/access"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var s = NewAccessHandlerTestSuite()

func TestAccessHandler_HandleGetCSRFToken(t *testing.T) {
	s.StartTest(t)

	s.ExpectUsecase().CreateTokenBySession(test.SessionID).Return(test.CSRFToken, nil)

	s.TestGetCSRFToken(s.Error("no session cookie"), http.StatusBadRequest)

	s.AddCookie(delivery.SessionIDCookieName, test.SessionID)
	s.TestGetCSRFToken(`{"status":"ok","body":{"token":"`+test.CSRFToken+`"}}`, http.StatusOK)
}

func TestAccessHandler_CsrfMiddleware(t *testing.T) {
	s.StartTest(t)

	next := func(c echo.Context) error {
		return c.String(http.StatusOK, "middleware passed")
	}
	middleware := s.handler.csrfMiddleware(next)

	s.SetupRequest(http.MethodGet, delivery.ApiV1UserListPath, "")
	middleware(s.NewContext())
	s.TestResponse("middleware passed", http.StatusOK)

	s.SetupRequest(http.MethodPost, delivery.ApiV1UserLogoutPath, "")
	middleware(s.NewContext())
	s.TestResponse(s.Error("csrf token required"), http.StatusBadRequest)

	s.ExpectUsecase().CheckTokenBySession(test.CSRFToken, test.SessionID).Return(true, nil)
	s.AddCookie(delivery.SessionIDCookieName, test.SessionID)
	s.SetupRequest(http.MethodPost, delivery.ApiV1UserLogoutPath, "")
	s.Request.Header.Add(csrfTokenHeader, test.CSRFToken)
	middleware(s.NewContext())
	s.TestResponse("middleware passed", http.StatusOK)

	s.ExpectUsecase().CheckTokenBySession(test.CSRFToken, test.SessionID).Return(false, nil)
	s.SetupRequest(http.MethodPost, delivery.ApiV1UserLogoutPath, "")
	s.Request.Header.Add(csrfTokenHeader, test.CSRFToken)
	middleware(s.NewContext())
	s.TestResponse(s.Error("incorrect CSRF token"), http.StatusBadRequest)
}

func TestCors(t *testing.T) {
	s.StartTest(t)

	middleware := s.handler.corsMiddleware(func(c echo.Context) error {
		return nil
	})

	s.SetupRequest(http.MethodOptions, "/", "")
	s.Request.Header.Add(echo.HeaderOrigin, "https://random324.now.sh")
	middleware(s.NewContext())
	s.TestResponse("", http.StatusNoContent)
	allowOrigin := s.Response.Header().Get(echo.HeaderAccessControlAllowOrigin)
	assert.Equal(t, allowOrigin, "", "Origin must be not allowed")

	s.SetupRequest(http.MethodOptions, "/", "")
	s.Request.Header.Add(echo.HeaderOrigin, "http://localhost:3000")
	middleware(s.NewContext())
	s.TestResponse("", http.StatusNoContent)
	allowOrigin = s.Response.Header().Get(echo.HeaderAccessControlAllowOrigin)
	assert.Equal(t, allowOrigin, "http://localhost:3000", "Origin must be allowed")
}

func TestIsOriginAllowed(t *testing.T) {
	assert.False(t, s.handler.isOriginAllowed(""))
	assert.True(t, s.handler.isOriginAllowed("https://20192lemmasnew-h0ivhao1a.now.sh"))
	assert.False(t, s.handler.isOriginAllowed("https://yandex.ru"))
	assert.False(t, s.handler.isOriginAllowed("https://random324.now.sh"))
	assert.True(t, s.handler.isOriginAllowed("http://localhost:3000"))
}

type AccessHandlerTestSuite struct {
	test.HandlerTestSuite
	usecase *access.MockCsrfUsecase
	handler *AccessHandler
}

func NewAccessHandlerTestSuite() *AccessHandlerTestSuite {
	return &AccessHandlerTestSuite{
		HandlerTestSuite: *test.NewHandlerTestSuite(),
	}
}

func (s *AccessHandlerTestSuite) StartTest(t *testing.T) {
	s.HandlerTestSuite.StartTest(t)
	s.usecase = access.NewMockCsrfUsecase(gomock.NewController(t))
	logger := testMock.NewMockLogger()
	s.handler = NewAccessHandler(s.E, s.usecase, logger)
}

func (s AccessHandlerTestSuite) ExpectUsecase() *access.MockCsrfUsecaseMockRecorder {
	return s.usecase.EXPECT()
}

func (s AccessHandlerTestSuite) TestGetCSRFToken(expectedResponse string, expectedCode int) {
	s.SetupRequestWithBody("")
	s.handler.handleGetCSRFToken(s.NewContext())
	s.TestResponse(expectedResponse, expectedCode)
}
