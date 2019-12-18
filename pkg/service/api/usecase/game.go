package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
	"io"
)

type gameUsecase struct {
	game game.GameClient
	log  logger.Logger

	c context.Context
}

func NewGameUsecase(game game.GameClient, log logger.Logger) api.GameUsecase {
	return &gameUsecase{
		game: game,
		log:  log,
		c:    context.Background(),
	}
}

func (u *gameUsecase) StartGame(userID int) error {
	res, err := u.game.StartGame(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *gameUsecase) StopGame(userID int) error {
	res, err := u.game.StopGame(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *gameUsecase) SetDirection(userID int, direction int) error {
	params := &game.UserAndDirection{UserId: int32(userID), Direction: int32(direction)}
	res, err := u.game.SetDirection(u.c, params)
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *gameUsecase) SetSpeed(userID int, speed int) error {
	params := &game.UserAndSpeed{UserId: int32(userID), Speed: int32(speed)}
	res, err := u.game.SetSpeed(u.c, params)
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *gameUsecase) GetPlayer(userID int) (*model.Player, error) {
	res, err := u.game.GetPlayer(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	return u.convertPlayer(res.Player), nil
}

func (u *gameUsecase) GetPlayers(userID int) ([]*model.Player, error) {
	res, err := u.game.GetPlayers(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	players := make([]*model.Player, 0)
	for _, p := range res.Players {
		players = append(players, u.convertPlayer(p))
	}
	return players, nil
}

func (u *gameUsecase) GetFood(userID int) ([]model.Food, error) {
	res, err := u.game.GetFood(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	food := make([]model.Food, 0)
	for _, f := range res.Food {
		food = append(food, u.convertFood(f))
	}
	return food, nil
}

func (u *gameUsecase) ListenEvents(userID int) (<-chan map[string]interface{}, error) {
	stream, err := u.game.ListenEvents(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	result := make(chan map[string]interface{})
	go func() {
		for {
			event, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				u.log.Error(err)
				return
			}
			result <- u.convertEvent(event)
		}
	}()
	return result, nil
}

func (u *gameUsecase) StopListenEvents(userID int) error {
	res, err := u.game.StopGame(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return err
	}
	if res.Error != "" {
		return errors.New(res.Error)
	}
	return nil
}

func (u *gameUsecase) convertUserID(id int) *game.UserID {
	return &game.UserID{UserId: int32(id)}
}

func (u *gameUsecase) convertPosition(pos *game.Position) model.Position {
	return model.Position{X: int(pos.X), Y: int(pos.Y)}
}

func (u *gameUsecase) convertPlayer(player *game.Player) *model.Player {
	return &model.Player{
		UserID:    int(player.UserId),
		Size:      int(player.Size),
		Direction: int(player.Direction),
		Speed:     int(player.Speed),
		Position:  u.convertPosition(player.Position),
	}
}

func (u *gameUsecase) convertFood(player *game.Food) model.Food {
	return model.Food{
		ID:       int(player.Id),
		Position: u.convertPosition(player.Position),
	}
}

func (u *gameUsecase) convertEvent(event *game.Event) map[string]interface{} {
	result := map[string]interface{}{}
	json.Unmarshal([]byte(event.Params), &result)
	return result
}
