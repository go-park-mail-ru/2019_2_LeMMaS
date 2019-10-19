package test

import (
	"encoding/json"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	"github.com/labstack/echo"
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

func (s HandlerTestSuite) TestResponse(expectedResponse string) {
	s.TestResponseStatus()
	s.TestResponseBody(expectedResponse)
	s.updateCookies()
}

func (s HandlerTestSuite) TestResponseBody(expectedResponse string) {
	actualBody := s.Response.Body.String()
	if strings.TrimSpace(expectedResponse) != strings.TrimSpace(actualBody) {
		s.T.Errorf("\nexpected response body:\n%v\ngot:\n%v", expectedResponse, actualBody)
	}
}

func (s HandlerTestSuite) TestResponseStatus() {
	if s.Response.Code != http.StatusOK {
		s.T.Errorf("expected status code 200, got %v", s.Response.Code)
	}
}

func (s HandlerTestSuite) TestCookiePresent(cookieName string) {
	if _, present := s.CookiesByName[cookieName]; !present {
		s.T.Errorf("cookie %v not found in response", cookieName)
	}
}

func (s HandlerTestSuite) TestCookieNotPresent(cookieName string) {
	if _, present := s.CookiesByName[cookieName]; present {
		s.T.Errorf("cookie %v must not be present in response", cookieName)
	}
}

func (s HandlerTestSuite) GetResponseBody() (map[string]*json.RawMessage, error) {
	var response map[string]*json.RawMessage
	err := json.NewDecoder(s.Response.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	var responseBody map[string]*json.RawMessage
	err = json.Unmarshal(*response["body"], &responseBody)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
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
