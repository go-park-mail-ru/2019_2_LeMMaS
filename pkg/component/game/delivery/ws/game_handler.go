package ws

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
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
}

func NewGameHandler(e *echo.Echo, gameUsecase game.Usecase, logger logger.Logger) *GameHandler {
	handler := GameHandler{logger: logger, gameUsecase: gameUsecase}
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	defer conn.Close()

	go func() {
		updates := h.gameUsecase.GetUpdatesStream()
		for update := range updates {
			h.OkWithBody(conn, update)
		}
	}()

	for {
		h.processRequest(conn)
	}
}

func (h GameHandler) processRequest(c *websocket.Conn) {
	request := wsDelivery.Request{}
	if err := c.ReadJSON(&request); err != nil {
		h.Error(c, "invalid json")
		return
	}
	switch request.Type {
	case "start":
		h.processGameStart(c)
	case "direction":
		h.processSetDirection(c)
	case "speed":
		h.processSetSpeed(c)
	default:
		h.Error(c, "unknown request type")
	}
}

type gameStartResponse struct {
	PlayerPosition model.Position   `json:"player_position"`
	Foods          []model.Position `json:"foods"`
}

func (h GameHandler) processGameStart(c *websocket.Conn) {
	h.OkWithBody(c, gameStartResponse{
		PlayerPosition: model.Position{10, 20},
		Foods:          []model.Position{{2, 1}, {432, 1}},
	})
}

type directionRequest struct {
	Direction int `json:"direction"`
}

func (h GameHandler) processSetDirection(c *websocket.Conn) {
	request := directionRequest{}
	if err := c.ReadJSON(&request); err != nil {
		h.Error(c, "invalid json")
		return
	}
	err := h.gameUsecase.SetDirection(request.Direction)
	if err != nil {
		h.Error(c, err.Error())
		return
	}
	h.Ok(c)
}

type speedRequest struct {
	Speed int `json:"speed"`
}

func (h GameHandler) processSetSpeed(c *websocket.Conn) {
	request := speedRequest{}
	if err := c.ReadJSON(&request); err != nil {
		h.Error(c, "invalid json")
		return
	}
	err := h.gameUsecase.SetSpeed(request.Speed)
	if err != nil {
		h.Error(c, err.Error())
		return
	}
	h.Ok(c)
}
