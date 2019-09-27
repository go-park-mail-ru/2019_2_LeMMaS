package handlers

import (
	"../config"
	"../cookie"
	"encoding/json"
	"net/http"
)

func MethodMiddleware(method string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				w.WriteHeader(405)
				w.Write([]byte("Method not allowed"))
			}

			next.ServeHTTP(w, r)
		})
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user config.AuthConfig
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Secure:          true,
		HTTPOnly:        true,
	}
	cookie.SetUserCookie(w, c, user.Login)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	curCookie, err := r.Cookie("sessionId")
	if err == http.ErrNoCookie || curCookie == nil {
		w.WriteHeader(403)
		return
	}
	cookie.DeleteCookie(w, r, *curCookie)
	w.WriteHeader(200)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user config.AuthConfig
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	if user.Login == "" || user.Email == "" { // TODO проверить наличие в БД
		w.WriteHeader(409)
		return
	}
	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Secure:          true,
		HTTPOnly:        true,
	}
	cookie.SetUserCookie(w, c, user.Login)
	w.WriteHeader(200)
}
