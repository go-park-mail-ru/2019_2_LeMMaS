package cookie

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/config"
	db "github.com/go-park-mail-ru/2019_2_LeMMaS/database"
	"net/http"
)

func makeHashCookie(login string) string {
	salt := "pdfnw;lsdvp"
	hash := 0
	for char := range login + salt {
		// TODO: update hash cookie
		hash = hash + char
	}
	return string(hash)
}

func CreateCookie(c config.SessionConfig, login string) (cookie *http.Cookie) {
	value := makeHashCookie(login)
	cookie = &http.Cookie{
		Name:     c.Name,
		Value:    value,
		MaxAge:   c.LifetimeSeconds,
		Path: 	  c.Path,
		Secure:   c.Secure,
		HttpOnly: c.HTTPOnly,
	}
	return
}

func SetUserCookie(w http.ResponseWriter, c config.SessionConfig, login string) {
	http.SetCookie(w, CreateCookie(c, login))
}

func DeleteCookie(w http.ResponseWriter, c http.Cookie) {
	c.MaxAge = -1
	http.SetCookie(w, &c)
}

func IsInDB(c http.Cookie) bool { //
	var nullUser db.User
	curUser := db.GetUserByCookie(c.Value)
	if curUser == nullUser {
		return false
	}
	return true
}
