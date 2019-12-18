package ws

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

var errInvalidJSON = errors.New("invalid json")
var errUnknownRequestType = errors.New("unknown request type")

type GameHandler struct {
	delivery.Handler
	log  logger.Logger
	game api.GameUsecase
	auth api.AuthUsecase
}

func NewGameHandler(e *echo.Echo, game api.GameUsecase, auth api.AuthUsecase, log logger.Logger) *GameHandler {
	handler := GameHandler{
		game: game,
		auth: auth,
		log:  log,
	}
	e.GET(delivery.ApiV1GamePath, handler.handleGame)
	return &handler
}

func (h *GameHandler) error(c *websocket.Conn, err error) error {
	body := map[string]interface{}{
		"type":    "error",
		"message": err.Error(),
	}
	return c.WriteJSON(body)
}

func (h *GameHandler) handleGame(c echo.Context) error {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		h.log.Error(err)
		return nil
	}
	defer conn.Close()

	userID, err := h.currentUserID(c)
	if err != nil {
		h.error(conn, err)
		return nil
	}

	for {
		err := h.processRequest(userID, conn)
		if err != nil {
			h.game.StopGame(userID)
			return nil
		}
	}
}

type request struct {
	Type string `json:"type"`
}

func (h *GameHandler) processRequest(userID int, c *websocket.Conn) error {
	_, requestReader, _ := c.NextReader()
	if requestReader == nil {
		return errors.New("no content")
	}
	requestBytes, err := ioutil.ReadAll(requestReader)
	if err != nil {
		return err
	}
	request := request{}
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return h.error(c, errInvalidJSON)
	}
	switch request.Type {
	case "start":
		return h.processGameStart(userID, c)
	case "stop":
		return h.processGameStop(userID, c)
	case "direction":
		return h.processSetDirection(userID, c, requestBytes)
	case "speed":
		return h.processSetSpeed(userID, c, requestBytes)
	default:
		return h.error(c, errUnknownRequestType)
	}
}

func (h *GameHandler) processGameStart(userID int, c *websocket.Conn) error {
	if err := h.game.StartGame(userID); err != nil {
		return h.error(c, err)
	}
	go func() {
		events, err := h.game.ListenEvents(userID)
		if err != nil {
			h.error(c, err)
			return
		}
		for event := range events {
			err := c.WriteJSON(event)
			if err != nil {
				h.game.StopListenEvents(userID)
				break
			}
		}
	}()
	players, err := h.game.GetPlayers(userID)
	if err != nil {
		return h.error(c, err)
	}
	food, err := h.game.GetFood(userID)
	if err != nil {
		return h.error(c, err)
	}
	return c.WriteJSON(map[string]interface{}{
		"type":    "start",
		"players": players,
		"food":    food,
	})
}

func (h *GameHandler) processGameStop(userID int, c *websocket.Conn) error {
	err := h.game.StopGame(userID)
	if err != nil {
		return h.error(c, err)
	}
	return nil
}

type directionRequest struct {
	Direction int `json:"direction"`
}

func (h *GameHandler) processSetDirection(userID int, c *websocket.Conn, data []byte) error {
	request := directionRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return h.error(c, errInvalidJSON)
	}
	err := h.game.SetDirection(userID, request.Direction)
	if err != nil {
		return h.error(c, err)
	}
	return nil
}

type speedRequest struct {
	Speed int `json:"speed"`
}

func (h *GameHandler) processSetSpeed(userID int, c *websocket.Conn, data []byte) error {
	request := speedRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return h.error(c, errInvalidJSON)
	}
	err := h.game.SetSpeed(userID, request.Speed)
	if err != nil {
		return h.error(c, err)
	}
	return nil
}

func (h *GameHandler) currentUserID(c echo.Context) (int, error) {
	cookie, err := c.Cookie(delivery.SessionCookieName)
	if err != nil {
		return 0, errors.New("no session cookie")
	}
	return h.auth.GetUserID(cookie.Value)
}
