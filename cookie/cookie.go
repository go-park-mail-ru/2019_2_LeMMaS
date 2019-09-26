package cookie

import (
	"net/http"
)

func CreateCookie(value string, c config.SessionConfig) (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:     c.Name,
		Value:    value,
		MaxAge:   c.LifetimeSeconds,
		HttpOnly: c.HTTPOnly,
		Path:     c.Path,
	}
	return
}

func GetSessionCookie(r *http.Request, c config.SessionConfig) (string, error) {
	session, err := r.Cookie(c.Name)
	if err != nil || session == nil {
		return "", err
	}
	return session.Value, err
}

func CreateAndSet(w http.ResponseWriter, c config.SessionConfig, value string) {
	http.SetCookie(w, CreateCookie(value, c))
}
