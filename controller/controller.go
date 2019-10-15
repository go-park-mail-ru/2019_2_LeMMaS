package controller

import (
	"encoding/json"
	"net/http"
	"time"
)

type Controller struct {
}

type APIResponse struct {
	Status string                 `json:"status"`
	Body   map[string]interface{} `json:"body"`
}

func (c Controller) writeCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (c Controller) writeOk(w http.ResponseWriter) {
	c.writeOkWithBody(w, map[string]interface{}{})
}

func (c Controller) writeOkWithBody(w http.ResponseWriter, body map[string]interface{}) {
	response, _ := json.Marshal(APIResponse{
		Status: "ok",
		Body:   body,
	})
	w.Write(response)
}

func (c Controller) writeError(w http.ResponseWriter, err error) {
	response, _ := json.Marshal(APIResponse{
		Status: "error",
		Body: map[string]interface{}{
			"message": err.Error(),
		},
	})
	w.Write(response)
}

func (c Controller) setCookie(w http.ResponseWriter, name, value string, expires time.Time) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
		Secure:  true,
	}
	http.SetCookie(w, cookie)
}

func (c Controller) deleteCookie(w http.ResponseWriter, cookie *http.Cookie) {
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}
