package ws

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

var errInvalidJSON = errors.New("invalid json")
var errUnknownRequestType = errors.New("unknown request type")

type GameHandler struct {
	logger      logger.Logger
	gameUsecase game.Usecase
	userUsecase user.Usecase

	upgrader websocket.Upgrader
}

func NewGameHandler(e *echo.Echo, gameUsecase game.Usecase, userUsecase user.Usecase, logger logger.Logger) *GameHandler {
	handler := GameHandler{
		logger:      logger,
		gameUsecase: gameUsecase,
		userUsecase: userUsecase,
		upgrader:    newWebsocketUpgrader(),
	}
	e.GET(httpDelivery.ApiV1GamePath, handler.handleGame)
	return &handler
}

func newWebsocketUpgrader() websocket.Upgrader {
	return websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
}

func (h GameHandler) sendError(c *websocket.Conn, err error) error {
	body := map[string]interface{}{
		"type":    game.EventError,
		"message": err.Error(),
	}
	return c.WriteJSON(body)
}

func (h GameHandler) handleGame(c echo.Context) error {
	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		h.logger.Error(err)
		return nil
	}
	defer conn.Close()

	pCurrentUser, err := h.getCurrentUser(c)
	if err != nil {
		h.sendError(conn, err)
		return nil
	}
	currentUser := *pCurrentUser
	userID := currentUser.ID

	for {
		err := h.processRequest(userID, conn)
		if err != nil {
			h.gameUsecase.StopGame(userID)
			return nil
		}
	}
}

type request struct {
	Type string `json:"type"`
}

func (h GameHandler) processRequest(userID int, c *websocket.Conn) error {
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
		return h.sendError(c, errInvalidJSON)
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
		return h.sendError(c, errUnknownRequestType)
	}
}

func (h GameHandler) processGameStart(userID int, c *websocket.Conn) error {
	if err := h.gameUsecase.StartGame(userID); err != nil {
		return h.sendError(c, err)
	}
	go func() {
		events, err := h.gameUsecase.ListenEvents(userID)
		if err != nil {
			h.sendError(c, err)
			return
		}
		for event := range events {
			err := c.WriteJSON(event)
			if err != nil {
				h.gameUsecase.StopListenEvents(userID)
				break
			}
		}
	}()
	return c.WriteJSON(map[string]interface{}{
		"type":    game.EventStart,
		"players": h.gameUsecase.GetPlayers(userID),
		"food":    h.gameUsecase.GetFood(userID),
	})
}

func (h GameHandler) processGameStop(userID int, c *websocket.Conn) error {
	err := h.gameUsecase.StopGame(userID)
	if err != nil {
		return h.sendError(c, err)
	}
	return nil
}

type directionRequest struct {
	Direction int `json:"direction"`
}

func (h GameHandler) processSetDirection(userID int, c *websocket.Conn, data []byte) error {
	request := directionRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return h.sendError(c, errInvalidJSON)
	}
	err := h.gameUsecase.SetDirection(userID, request.Direction)
	if err != nil {
		return h.sendError(c, err)
	}
	return nil
}

type speedRequest struct {
	Speed int `json:"speed"`
}

func (h GameHandler) processSetSpeed(userID int, c *websocket.Conn, data []byte) error {
	request := speedRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return h.sendError(c, errInvalidJSON)
	}
	err := h.gameUsecase.SetSpeed(userID, request.Speed)
	if err != nil {
		return h.sendError(c, err)
	}
	return nil
}

func (h GameHandler) getCurrentUser(c echo.Context) (*model.User, error) {
	sessionIDCookie, err := c.Cookie(httpDelivery.SessionIDCookieName)
	if err != nil {
		return nil, errors.New("no session cookie")
	}
	currentUser, _ := h.userUsecase.GetUserBySessionID(sessionIDCookie.Value)
	if currentUser == nil {
		return nil, errors.New("invalid session id")
	}
	return currentUser, nil
}
