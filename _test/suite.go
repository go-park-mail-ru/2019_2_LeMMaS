package _test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type ControllerTestSuite struct {
	T             *testing.T
	Request       *http.Request
	Response      *httptest.ResponseRecorder
	CookiesByName map[string]*http.Cookie
}

func NewControllerTestSuite() *ControllerTestSuite {
	return &ControllerTestSuite{
		CookiesByName: map[string]*http.Cookie{},
	}
}

func (s *ControllerTestSuite) SetTesting(t *testing.T) {
	s.T = t
}

func (s ControllerTestSuite) TestResponse(expectedResponse string) {
	s.TestResponseStatus()
	s.TestResponseBody(expectedResponse)
	s.updateCookies()
}

func (s ControllerTestSuite) TestResponseBody(expectedResponse string) {
	actualBody := s.Response.Body.String()
	if strings.TrimSpace(expectedResponse) != strings.TrimSpace(actualBody) {
		s.T.Errorf("\nexpected response body:\n%v\ngot:\n%v", expectedResponse, actualBody)
	}
}

func (s ControllerTestSuite) TestResponseStatus() {
	if s.Response.Code != http.StatusOK {
		s.T.Errorf("expected status code 200, got %v", s.Response.Code)
	}
}

func (s ControllerTestSuite) TestCookiePresent(cookieName string) {
	if _, present := s.CookiesByName[cookieName]; !present {
		s.T.Errorf("cookie %v not found in response", cookieName)
	}
}

func (s ControllerTestSuite) TestCookieNotPresent(cookieName string) {
	if _, present := s.CookiesByName[cookieName]; present {
		s.T.Errorf("cookie %v must not be present in response", cookieName)
	}
}

func (s ControllerTestSuite) GetResponseBody() (map[string]*json.RawMessage, error) {
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

func (s ControllerTestSuite) updateCookies() {
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
