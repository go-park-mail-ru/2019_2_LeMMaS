package ws

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	wsDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/ws"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"net/http"
)

type GameHandler struct {
	wsDelivery.Handler
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
		h.Error(conn, err.Error())
		return nil
	}

	go func() {
		updates := h.gameUsecase.GetEventStream(*currUser)
		for update := range updates {
			h.OkWithBody(conn, update)
		}
	}()

	for {
		err := h.processRequest(*currUser, conn)
		if err != nil {
			return nil
		}
	}
}

func (h GameHandler) processRequest(user model.User, conn *websocket.Conn) error {
	request := wsDelivery.Request{}
	err := conn.ReadJSON(&request)
	if err != nil {
		return err
	}
	switch request.Type {
	case "start":
		return h.processGameStart(user, conn)
	case "direction":
		return h.processSetDirection(user, conn)
	case "speed":
		return h.processSetSpeed(user, conn)
	default:
		return h.Error(conn, "unknown request type")
	}
}

type gameStartResponse struct {
	PlayerPosition model.Position   `json:"player_position"`
	Foods          []model.Position `json:"foods"`
}

func (h GameHandler) processGameStart(user model.User, c *websocket.Conn) error {
	h.gameUsecase.StartGame(user)
	return h.OkWithBody(c, gameStartResponse{
		PlayerPosition: h.gameUsecase.GetPlayerPosition(user),
		Foods:          h.gameUsecase.GetFoodsPositions(user),
	})
}

type directionRequest struct {
	Direction float64 `json:"direction"`
}

func (h GameHandler) processSetDirection(user model.User, conn *websocket.Conn) error {
	request := directionRequest{}
	if err := conn.ReadJSON(&request); err != nil {
		return h.Error(conn, "invalid json")
	}
	err := h.gameUsecase.SetDirection(user, request.Direction)
	if err != nil {
		return h.Error(conn, err.Error())
	}
	return h.Ok(conn)
}

type speedRequest struct {
	Speed float64 `json:"speed"`
}

func (h GameHandler) processSetSpeed(user model.User, c *websocket.Conn) error {
	request := speedRequest{}
	if err := c.ReadJSON(&request); err != nil {
		return h.Error(c, "invalid json")
	}
	err := h.gameUsecase.SetSpeed(user, request.Speed)
	if err != nil {
		return h.Error(c, err.Error())
	}
	return h.Ok(c)
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
