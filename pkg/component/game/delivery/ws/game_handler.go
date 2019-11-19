package ws

import (
	"encoding/json"
	"fmt"
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

type GameHandler struct {
	logger      logger.Logger
	gameUsecase game.Usecase
	userUsecase user.Usecase
}

func NewGameHandler(e *echo.Echo, gameUsecase game.Usecase, userUsecase user.Usecase, logger logger.Logger) *GameHandler {
	handler := GameHandler{
		logger:      logger,
		gameUsecase: gameUsecase,
		userUsecase: userUsecase,
	}
	e.GET(httpDelivery.ApiV1GamePath, handler.HandleGame)
	return &handler
}

func (h GameHandler) sendEvent(c *websocket.Conn, event model.GameEvent) error {
	return c.WriteJSON(event)
}

func (h GameHandler) sendError(c *websocket.Conn, message string) error {
	body := map[string]interface{}{
		"type":    model.GameEventError,
		"message": message,
	}
	return c.WriteJSON(body)
}

func (h GameHandler) HandleGame(c echo.Context) error {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		h.logger.Error(err)
		return nil
	}
	defer conn.Close()

	currUser, err := h.getCurrentUser(c)
	if err != nil {
		h.sendError(conn, err.Error())
		return nil
	}

	go func() {
		events := h.gameUsecase.GetEventsStream(*currUser)
		for event := range events {
			h.sendEvent(conn, event)
		}
	}()

	for {
		err := h.processRequest(*currUser, conn)
		if err != nil {
			return nil
		}
	}
}

type Request struct {
	Type string `json:"type"`
}

func (h GameHandler) processRequest(user model.User, c *websocket.Conn) error {
	_, requestReader, _ := c.NextReader()
	requestBytes, err := ioutil.ReadAll(requestReader)
	if err != nil {
		return err
	}
	request := Request{}
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return h.sendError(c, "invalid json")
	}
	switch request.Type {
	case "start":
		return h.processGameStart(user, c)
	case "direction":
		return h.processSetDirection(user, c, requestBytes)
	case "speed":
		return h.processSetSpeed(user, c, requestBytes)
	default:
		return h.sendError(c, "unknown request type")
	}
}

func (h GameHandler) processGameStart(user model.User, c *websocket.Conn) error {
	h.gameUsecase.StartGame(user)
	return h.sendEvent(c, map[string]interface{}{
		"type":    model.GameEventStart,
		"players": h.convertPlayersToOutput(h.gameUsecase.GetPlayers(user)),
		"foods":   h.convertFoodToOutput(h.gameUsecase.GetFood(user)),
	})
}

func (h GameHandler) convertPlayersToOutput(playersByID map[int]*model.Player) []map[string]interface{} {
	players := make([]map[string]interface{}, 0, len(playersByID))
	for id, player := range playersByID {
		players = append(players, map[string]interface{}{
			"id": id,
			"x":  player.Position.X,
			"y":  player.Position.Y,
		})
	}
	return players
}

func (h GameHandler) convertFoodToOutput(foodByID map[int]*model.Position) []map[string]interface{} {
	foods := make([]map[string]interface{}, 0, len(foodByID))
	for id, food := range foodByID {
		foods = append(foods, map[string]interface{}{
			"id": id,
			"x":  food.X,
			"y":  food.Y,
		})
	}
	return foods
}

type directionRequest struct {
	Direction int `json:"direction"`
}

func (h GameHandler) processSetDirection(user model.User, c *websocket.Conn, data []byte) error {
	request := directionRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return h.sendError(c, "invalid json")
	}
	err := h.gameUsecase.SetDirection(user, request.Direction)
	if err != nil {
		return h.sendError(c, err.Error())
	}
	return nil
}

type speedRequest struct {
	Speed int `json:"speed"`
}

func (h GameHandler) processSetSpeed(user model.User, c *websocket.Conn, data []byte) error {
	request := speedRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return h.sendError(c, "invalid json")
	}
	err := h.gameUsecase.SetSpeed(user, request.Speed)
	if err != nil {
		return h.sendError(c, err.Error())
	}
	return nil
}

func (h GameHandler) getCurrentUser(c echo.Context) (*model.User, error) {
	sessionIDCookie, err := c.Cookie(httpDelivery.SessionIDCookieName)
	if err != nil {
		return nil, fmt.Errorf("no session cookie")
	}
	currentUser, _ := h.userUsecase.GetUserBySessionID(sessionIDCookie.Value)
	if currentUser == nil {
		return nil, fmt.Errorf("invalid session id")
	}
	return currentUser, nil
}
