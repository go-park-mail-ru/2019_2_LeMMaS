package test

import (
	"encoding/json"
	"fmt"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type HandlerTestSuite struct {
	T             *testing.T
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

func (s *HandlerTestSuite) SetTesting(t *testing.T) {
	s.T = t
	s.E = echo.New()
}

func (s *HandlerTestSuite) SetupRequest(method, path string, requestBody string) {
	s.Response = httptest.NewRecorder()
	s.Request = httptest.NewRequest(method, path, strings.NewReader(requestBody))
	s.Request.Header.Add("Content-Type", "application/json")
	for _, cookie := range s.CookiesByName {
		s.Request.AddCookie(cookie)
	}
}

func (s *HandlerTestSuite) NewHandlerContext() echo.Context {
	return s.E.NewContext(s.Request, s.Response)
}

func (s HandlerTestSuite) TestOkResponse(expectedResponse string) {
	assert.Equal(s.T, s.Response.Code, http.StatusOK)
	s.TestResponseBody(expectedResponse)
	s.updateCookies()
}

func (s HandlerTestSuite) TestResponse(expectedResponse string, expectedCode int) {
	assert.Equal(s.T, s.Response.Code, expectedCode)
	s.TestResponseBody(expectedResponse)
	s.updateCookies()
}

func (s HandlerTestSuite) TestResponseBody(expectedResponse string) {
	actualBody := s.Response.Body.String()
	assert.Equal(s.T, strings.TrimSpace(expectedResponse), strings.TrimSpace(actualBody), "unexpected response body")
}

func (s HandlerTestSuite) TestCookiePresent(cookieName string) {
	_, present := s.CookiesByName[cookieName]
	assert.True(s.T, present, fmt.Sprintf("cookie %v not found in response", cookieName))
}

func (s HandlerTestSuite) TestCookieNotPresent(cookieName string) {
	_, present := s.CookiesByName[cookieName]
	assert.False(s.T, present, fmt.Sprintf("cookie %v must not be present in response", cookieName))
}

func (s *HandlerTestSuite) Ok() string {
	response := httpDelivery.Response{
		Status: "ok",
	}
	result, _ := json.Marshal(response)
	return string(result)
}

func (s *HandlerTestSuite) Error(message string) string {
	response := httpDelivery.Response{
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