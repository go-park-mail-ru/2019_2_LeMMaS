 package handlers

 import (
	 "github.com/go-park-mail-ru/2019_2_LeMMaS/config"
	 "github.com/go-park-mail-ru/2019_2_LeMMaS/cookie"
	 db "github.com/go-park-mail-ru/2019_2_LeMMaS/database"
	 "io"
	 "net/http"
	 "net/http/httptest"
	 "strings"
	 "testing"
 )

 var (
	 users []db.User
 )

type testRW struct {
	Cookie       http.Cookie
	Body         io.Reader
	Response     string
	StatusCode   int
}

func initDB () {
	var curUser = config.AuthConfig{
		Login:    "Pavel",
		Password: "password",
		Email:    "pavel@mail.ru",
	}
	_ = db.CreateNewUser(curUser)
}

func TestLoginHandler(t *testing.T) {
	initDB()

	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	curCookie := cookie.CreateCookie(c, "Pavel")

	cases := []testRW {
		{
			Response:     `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       *curCookie,
			Response:     `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       http.Cookie{Name: "sessionId", Value: "no"}, // левые куки
			Response:     `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Body:         strings.NewReader(`{"login":"Pavel","password":"password"}`),
			StatusCode:   http.StatusOK, // 200
		},
		{
			Cookie:       *curCookie, // уже с куками
			Body:         strings.NewReader(`{"login":"Pavel","password":"password"}`),
			Response:     `"error":"already logged in"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Body:         strings.NewReader(`{"login":"Pavel","password":"no"}`), // левый пароль
			Response:     `"error":"invalid credentials"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
	}

	for i, item := range cases {
		url := "api/v1/login"
		r := httptest.NewRequest("POST", url, item.Body)
		r.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		LoginHandler(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("%d) wrong StatusCode: got %d, expected %d", i, w.Code, item.StatusCode)
		}
	}
}


func TestLogoutHandler(t *testing.T) {
	initDB()

	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	curCookie := cookie.CreateCookie(c, "Pavel")

	cases := []testRW {
		{
			Response:     `"error":"invalid credentials"`, // без кук
			StatusCode:   http.StatusForbidden, // 403
		},
		{
			Cookie:       *curCookie,
			StatusCode:   http.StatusOK, // 200
		},
		{
			Cookie:       http.Cookie{Name: "sessionId", Value: "no"}, // левые куки
			Response:     `"error":"invalid credentials"`,
			StatusCode:   http.StatusForbidden, // 403
		},
	}

	for i, item := range cases {
		url := "api/v1/logout"
		r := httptest.NewRequest("POST", url, item.Body)
		r.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		LogoutHandler(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("%d) wrong StatusCode: got %d, expected %d", i, w.Code, item.StatusCode)
		}
	}
}


func TestRegisterHandler(t *testing.T) {
	initDB()

	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	curCookie := cookie.CreateCookie(c, "Pavel")

	cases := []testRW {
		{
			Response:     `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Body:         strings.NewReader(`{"login":"Pavel1","password":"password","email":"pava@mail.ru"}`),
			StatusCode:   http.StatusOK, // 200
		},
		{
			Cookie:       *curCookie, // уже с куками
			Response:     `"error":"already logged in"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       *curCookie, // уже с куками
			Body:         strings.NewReader(`{"login":"Pavel","password":"password"}`),
			Response:     `"error":"already logged in"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Body:         strings.NewReader(`{"login":"Pavel","password":"no","email":"pava@mail.ru"}`), // уже есть логин
			Response:     `"error":"login already exists"`,
			StatusCode:   http.StatusConflict, // 409
		},
		{
			Body:         strings.NewReader(`{"login":"Vasya","password":"no","email":"pavel@mail.ru"}`), // уже есть email
			Response:     `"error":"email already exists"`,
			StatusCode:   http.StatusConflict, // 409
		},
	}

	for i, item := range cases {
		url := "api/v1/register"
		r := httptest.NewRequest("POST", url, item.Body)
		w := httptest.NewRecorder()

		RegisterHandler(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("%d) wrong StatusCode: got %d, expected %d", i, w.Code, item.StatusCode)
		}
	}
}


 func TestGetUserDataHandler(t *testing.T) {
	 initDB()

	 c := config.SessionConfig{
		 Name:            "sessionId",
		 LifetimeSeconds: 86400,
		 Path:			 "/",
		 Secure:          true,
		 HTTPOnly:        true,
	 }
	 curCookie := cookie.CreateCookie(c, "Pavel")

	 cases := []testRW {
		 {
			 Response:     `"error":"invalid credentials"`, // без кук
			 StatusCode:   http.StatusUnauthorized, // 401
		 },
		 {
			 Cookie:       *curCookie, // с куками
			 StatusCode:   http.StatusOK, // 200
		 },
		 {
			 Cookie:       http.Cookie{Name: "sessionId", Value: "no"}, // левые куки
			 Response:     `"error":"invalid credentials"`,
			 StatusCode:   http.StatusForbidden, // 403
		 },
	 }

	 for i, item := range cases {
		 url := "api/v1/user"
		 r := httptest.NewRequest("GET", url, item.Body)
		 w := httptest.NewRecorder()
		 r.AddCookie(&item.Cookie)

		 GetUserDataHandler(w, r)

		 if w.Code != item.StatusCode {
			 t.Errorf("%d) wrong StatusCode: got %d, expected %d", i, w.Code, item.StatusCode)
		 }
	 }
 }

func TestChangeUserDataHandler(t *testing.T) {
	initDB()

	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	curCookie := cookie.CreateCookie(c, "Pavel")

	cases := []testRW {
		{
			Response:     `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       *curCookie, // с куками
			Response:     `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       *curCookie, // с куками
			Body:         strings.NewReader(`{"login":"Pasha","password":"no","email":"pava@mail.ru"}`),
			StatusCode:   http.StatusOK, // 200
		},
		{
			Cookie:       *curCookie, // с куками
			Body:         strings.NewReader(`{"email":"pava@mail.ru"}`),
			StatusCode:   http.StatusOK, // 200
		},
		{
			Cookie:       http.Cookie{Name: "sessionId", Value: "no"}, // левые куки
			Response:     `"error":"invalid credentials"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       http.Cookie{Name: "sessionId", Value: "no"}, // левые куки
			Body:         strings.NewReader(`{"login":"Pasha","password":"no","email":"pava@mail.ru"}`),
			Response:     `"error":"invalid credentials"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
	}

	for i, item := range cases {
		url := "api/v1/settings"
		r := httptest.NewRequest("PATCH", url, item.Body)
		w := httptest.NewRecorder()
		r.AddCookie(&item.Cookie)

		ChangeUserDataHandler(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("%d) wrong StatusCode: got %d, expected %d", i, w.Code, item.StatusCode)
		}
	}
}


func TestUploadAvatarHandler(t *testing.T) {
	initDB()

	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	curCookie := cookie.CreateCookie(c, "Pavel")

	cases := []testRW {
		{
			Response:     `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       *curCookie, // с куками
			Response:     `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       http.Cookie{Name: "sessionId", Value: "no"}, // левые куки
			Response:     `"error":"invalid credentials"`,
			StatusCode:   http.StatusBadRequest, // 400
		},
		{
			Cookie:       *curCookie, // с куками
			// TODO здесь надо как-то считывать данные формы
			StatusCode:   http.StatusOK, // 200
		},
	}

	for i, item := range cases {
		url := "api/v1/upload"
		r := httptest.NewRequest("PUT", url, item.Body)
		w := httptest.NewRecorder()
		r.AddCookie(&item.Cookie)

		UploadAvatarHandler(w, r)

		if w.Code != item.StatusCode {
			t.Errorf("%d) wrong StatusCode: got %d, expected %d", i, w.Code, item.StatusCode)
		}
	}
}