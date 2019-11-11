package http

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Handler struct {
}

type Response struct {
	Status string      `json:"status"`
	Body   interface{} `json:"body"`
}

func (h *Handler) Ok(c echo.Context) error {
	response := Response{
		Status: "ok",
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) OkWithBody(c echo.Context, body interface{}) error {
	response := Response{
		Status: "ok",
		Body:   body,
	}
	return c.JSON(http.StatusOK, response)
}

type errorResponseBody struct {
	Message string `json:"message"`
}

func (h *Handler) Error(c echo.Context, message string) error {
	response := Response{
		Status: "error",
		Body:   errorResponseBody{Message: message},
	}
	return c.JSON(http.StatusBadRequest, response)
}

func (h *Handler) Errors(c echo.Context, errors []error) error {
	message := ""
	for _, err := range errors {
		message += err.Error() + "   "
	}
	return h.Error(c, message)
}

func (h *Handler) Validate(data interface{}) (bool, []error) {
	ok, err := valid.ValidateStruct(data)
	if !ok {
		if errors, ok := err.(valid.Errors); ok {
			return false, errors
		}
	}
	return true, []error{}
}

func (h *Handler) SetCookie(c echo.Context, name, value string, expires time.Time) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
		Path:    "/",
	}
	c.SetCookie(cookie)
}

func (h *Handler) DeleteCookie(c echo.Context, name string) {
	yesterday := time.Now().AddDate(0, 0, -1)
	h.SetCookie(c, name, "", yesterday)
}
