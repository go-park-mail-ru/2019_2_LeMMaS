package ws

import (
	"github.com/gorilla/websocket"
)

type Handler struct {
}

type Request struct {
	Type string `json:"type"`
}

type Response struct {
	Status string      `json:"status"`
	Body   interface{} `json:"body"`
}

func (h *Handler) Ok(c *websocket.Conn) error {
	return c.WriteJSON(Response{Status: "ok"})
}

func (h *Handler) OkWithBody(c *websocket.Conn, body interface{}) error {
	response := Response{
		Status: "ok",
		Body:   body,
	}
	return c.WriteJSON(response)
}

type errorResponseBody struct {
	Message string `json:"message"`
}

func (h Handler) Error(c *websocket.Conn, message string) error {
	response := Response{
		Status: "error",
		Body:   errorResponseBody{Message: message},
	}
	return c.WriteJSON(response)
}
