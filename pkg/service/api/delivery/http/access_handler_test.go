package http

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var as = NewAccessHandlerTestSuite()

func TestAccessHandler_HandleGetCSRFToken(t *testing.T) {
	as.StartTest(t)

	as.Expect().CreateTokenBySession(test.SessionID).Return(test.CSRFToken, nil)

	as.TestGetCSRFToken(as.Error("no session cookie"), http.StatusBadRequest)

	as.AddCookie(delivery.SessionCookieName, test.SessionID)
	as.TestGetCSRFToken(`{"status":"ok","body":{"token":"`+test.CSRFToken+`"}}`, http.StatusOK)
}

func TestAccessHandler_CsrfMiddleware(t *testing.T) {
	as.StartTest(t)

	next := func(c echo.Context) error {
		return c.String(http.StatusOK, "middleware passed")
	}
	middleware := as.handler.csrfMiddleware(next)

	as.SetupRequest(http.MethodGet, delivery.ApiV1UserListPath, "")
	middleware(as.NewContext())
	as.TestResponse("middleware passed", http.StatusOK)

	as.SetupRequest(http.MethodPost, delivery.ApiV1UserLogoutPath, "")
	middleware(as.NewContext())
	as.TestResponse(as.Error("csrf token required"), http.StatusBadRequest)

	as.Expect().CheckTokenBySession(test.CSRFToken, test.SessionID).Return(true, nil)
	as.AddCookie(delivery.SessionCookieName, test.SessionID)
	as.SetupRequest(http.MethodPost, delivery.ApiV1UserLogoutPath, "")
	as.Request.Header.Add(csrfTokenHeader, test.CSRFToken)
	middleware(as.NewContext())
	as.TestResponse("middleware passed", http.StatusOK)

	as.Expect().CheckTokenBySession(test.CSRFToken, test.SessionID).Return(false, nil)
	as.SetupRequest(http.MethodPost, delivery.ApiV1UserLogoutPath, "")
	as.Request.Header.Add(csrfTokenHeader, test.CSRFToken)
	middleware(as.NewContext())
	as.TestResponse(as.Error("incorrect CSRF token"), http.StatusBadRequest)
}

func TestCors(t *testing.T) {
	as.StartTest(t)

	middleware := as.handler.corsMiddleware(func(c echo.Context) error {
		return nil
	})

	as.SetupRequest(http.MethodOptions, "/", "")
	as.Request.Header.Add(echo.HeaderOrigin, "https://random324.now.sh")
	middleware(as.NewContext())
	as.TestResponse("", http.StatusNoContent)
	allowOrigin := as.Response.Header().Get(echo.HeaderAccessControlAllowOrigin)
	assert.Equal(t, allowOrigin, "", "Origin must be not allowed")

	as.SetupRequest(http.MethodOptions, "/", "")
	as.Request.Header.Add(echo.HeaderOrigin, "http://localhost:3000")
	middleware(as.NewContext())
	as.TestResponse("", http.StatusNoContent)
	allowOrigin = as.Response.Header().Get(echo.HeaderAccessControlAllowOrigin)
	assert.Equal(t, allowOrigin, "http://localhost:3000", "Origin must be allowed")
}

func TestIsOriginAllowed(t *testing.T) {
	assert.False(t, as.handler.isOriginAllowed(""))
	assert.True(t, as.handler.isOriginAllowed("https://20192lemmasnew-h0ivhao1a.now.sh"))
	assert.False(t, as.handler.isOriginAllowed("https://yandex.ru"))
	assert.False(t, as.handler.isOriginAllowed("https://random324.now.sh"))
	assert.True(t, as.handler.isOriginAllowed("http://localhost:3000"))
}

type AccessHandlerTestSuite struct {
	test.HandlerTestSuite
	csrf    *api.MockCsrfUsecase
	handler *AccessHandler
}

func NewAccessHandlerTestSuite() *AccessHandlerTestSuite {
	return &AccessHandlerTestSuite{
		HandlerTestSuite: *test.NewHandlerTestSuite(),
	}
}

func (s *AccessHandlerTestSuite) StartTest(t *testing.T) {
	s.HandlerTestSuite.StartTest(t)
	s.csrf = api.NewMockCsrfUsecase(gomock.NewController(t))
	log := testMock.NewMockLogger(t)
	s.handler = NewAccessHandler(s.E, s.csrf, log)
}

func (s AccessHandlerTestSuite) Expect() *api.MockCsrfUsecaseMockRecorder {
	return s.csrf.EXPECT()
}

func (s AccessHandlerTestSuite) TestGetCSRFToken(expectedResponse string, expectedCode int) {
	s.SetupRequestWithBody("")
	s.handler.handleGetCSRFToken(s.NewContext())
	s.TestResponse(expectedResponse, expectedCode)
}
