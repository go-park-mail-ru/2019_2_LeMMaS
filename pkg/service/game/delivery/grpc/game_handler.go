package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type GameHandler struct {
	usecase game.Usecase
	server  *grpc.Server
	log     logger.Logger
}

func NewGameHandler(usecase game.Usecase, log logger.Logger) *GameHandler {
	h := GameHandler{usecase: usecase, log: log}
	h.server = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	game.RegisterGameServer(h.server, &h)
	return &h
}

func (h *GameHandler) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Println("listening " + address)
	return h.server.Serve(listener)
}

func (h *GameHandler) StartGame(c context.Context, params *game.UserID) (*game.Error, error) {
	res := &game.Error{}
	err := h.usecase.StartGame(int(params.UserId))
	if err != nil {
		res.Error = err.Error()
	}
	return res, nil
}

func (h *GameHandler) StopGame(c context.Context, params *game.UserID) (*game.Error, error) {
	res := &game.Error{}
	err := h.usecase.StopGame(int(params.UserId))
	if err != nil {
		res.Error = err.Error()
	}
	return res, nil
}

func (h *GameHandler) SetDirection(c context.Context, params *game.UserAndDirection) (*game.Error, error) {
	res := &game.Error{}
	err := h.usecase.SetDirection(int(params.UserId), int(params.Direction))
	if err != nil {
		res.Error = err.Error()
	}
	return res, nil
}

func (h *GameHandler) SetSpeed(c context.Context, params *game.UserAndSpeed) (*game.Error, error) {
	res := &game.Error{}
	err := h.usecase.SetSpeed(int(params.UserId), int(params.Speed))
	if err != nil {
		res.Error = err.Error()
	}
	return res, nil
}

func (h *GameHandler) GetPlayer(c context.Context, params *game.UserID) (*game.GetPlayerResult, error) {
	res := &game.GetPlayerResult{}
	player := h.usecase.GetPlayer(int(params.UserId))
	if player != nil {
		res.Player = h.convertPlayer(player)
	}
	return res, nil
}

func (h *GameHandler) GetPlayers(c context.Context, params *game.UserID) (*game.GetPlayersResult, error) {
	res := &game.GetPlayersResult{
		Players: make([]*game.Player, 0),
	}
	players := h.usecase.GetPlayers(int(params.UserId))
	for _, p := range players {
		res.Players = append(res.Players, h.convertPlayer(p))
	}
	return res, nil
}

func (h *GameHandler) GetFood(c context.Context, params *game.UserID) (*game.GetFoodResult, error) {
	res := &game.GetFoodResult{
		Food: make([]*game.Food, 0),
	}
	food := h.usecase.GetFood(int(params.UserId))
	for _, f := range food {
		res.Food = append(res.Food, h.convertFood(f))
	}
	return res, nil
}

func (h *GameHandler) ListenEvents(params *game.UserID, stream game.Game_ListenEventsServer) error {
	events, err := h.usecase.ListenEvents(int(params.UserId))
	if err != nil {
		err = stream.Send(&game.Event{Error: err.Error()})
		return err
	}
	for event := range events {
		err := stream.Send(h.convertEvent(event))
		if err != nil {
			h.log.Error(err)
			return err
		}
	}
	return nil
}

func (h *GameHandler) StopListenEvents(ctx context.Context, params *game.UserID) (result *game.Error, grpcErr error) {
	result = &game.Error{}
	err := h.usecase.StopListenEvents(int(params.UserId))
	if err != nil {
		result.Error = err.Error()
	}
	return
}

func (h *GameHandler) convertPlayer(player *model.Player) *game.Player {
	return &game.Player{
		UserId:    int32(player.UserID),
		Speed:     int32(player.Speed),
		Direction: int32(player.Direction),
		Size:      int32(player.Size),
		Position:  h.convertPosition(player.Position),
	}
}

func (h *GameHandler) convertFood(food model.Food) *game.Food {
	return &game.Food{
		Id:       int32(food.ID),
		Position: h.convertPosition(food.Position),
	}
}

func (h *GameHandler) convertPosition(pos model.Position) *game.Position {
	return &game.Position{X: int32(pos.X), Y: int32(pos.Y)}
}

func (h *GameHandler) convertEvent(event map[string]interface{}) *game.Event {
	params, _ := json.Marshal(event)
	return &game.Event{Params: string(params)}
}
