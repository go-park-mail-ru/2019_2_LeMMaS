package test

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const cookieExpiration = time.Hour * 10

type HandlerTestSuite struct {
	t             *testing.T
	E             *echo.Echo
	Request       *http.Request
	Response      *httptest.ResponseRecorder
	CookiesByName map[string]*http.Cookie
}

func NewHandlerTestSuite() *HandlerTestSuite {
	return &HandlerTestSuite{
		CookiesByName: map[string]*http.Cookie{},
	}
}

func (s *HandlerTestSuite) StartTest(t *testing.T) {
	s.t = t
	s.E = echo.New()
}

func (s *HandlerTestSuite) SetupRequestWithBody(requestBody string) {
	// The methods and path does not matter, if handler is called directly (as in test suite)
	s.SetupRequest(http.MethodGet, "/", requestBody)
}

func (s *HandlerTestSuite) SetupRequest(method, path, requestBody string) {
	s.Response = httptest.NewRecorder()
	s.Request = httptest.NewRequest(method, path, strings.NewReader(requestBody))
	s.Request.Header.Add("Content-Type", "application/json")
	for _, cookie := range s.CookiesByName {
		s.Request.AddCookie(cookie)
	}
}

func (s HandlerTestSuite) AddCookie(name, value string) {
	s.CookiesByName[name] = &http.Cookie{Name: name, Value: value, Expires: time.Now().Add(cookieExpiration)}
}

func (s *HandlerTestSuite) NewContext() echo.Context {
	return s.E.NewContext(s.Request, s.Response)
}

func (s HandlerTestSuite) TestOkResponse(expectedResponse string) {
	assert.Equal(s.t, http.StatusOK, s.Response.Code, "unexpected response status")
	s.testResponseBody(expectedResponse)
	s.updateCookies()
}

func (s HandlerTestSuite) TestResponse(expectedResponse string, expectedCode int) {
	assert.Equal(s.t, expectedCode, s.Response.Code, "unexpected response status")
	s.testResponseBody(expectedResponse)
	s.updateCookies()
}

func (s HandlerTestSuite) testResponseBody(expectedResponse string) {
	actualBody := s.Response.Body.String()
	assert.Equal(s.t, strings.TrimSpace(expectedResponse), strings.TrimSpace(actualBody), "unexpected response body")
}

func (s HandlerTestSuite) TestCookiePresent(cookieName string) {
	_, present := s.CookiesByName[cookieName]
	assert.True(s.t, present, fmt.Sprintf("cookie %v not found in response", cookieName))
}

func (s HandlerTestSuite) TestCookieNotPresent(cookieName string) {
	_, present := s.CookiesByName[cookieName]
	assert.False(s.t, present, fmt.Sprintf("cookie %v must not be present in response", cookieName))
}

func (s *HandlerTestSuite) Ok() string {
	response := delivery.Response{
		Status: "ok",
	}
	result, _ := json.Marshal(response)
	return string(result)
}

func (s *HandlerTestSuite) Error(message string) string {
	response := delivery.Response{
		Status: "error",
		Body: map[string]interface{}{
			"message": message,
		},
	}
	result, _ := json.Marshal(response)
	return string(result)
}

func (s HandlerTestSuite) updateCookies() {
	for _, cookie := range s.Response.Result().Cookies() {
		if cookie.Expires.Before(time.Now()) {
			if _, ok := s.CookiesByName[cookie.Name]; ok {
				delete(s.CookiesByName, cookie.Name)
			}
		} else {
			s.CookiesByName[cookie.Name] = cookie
		}
	}
}
