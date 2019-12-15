package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
)

type GameUsecase struct {
	game game.GameClient
	log  logger.Logger

	c context.Context
}

func NewGameUsecase(game game.GameClient, log logger.Logger) api.GameUsecase {
	return &GameUsecase{
		game: game,
		log:  log,
		c:    context.Background(),
	}
}

func (u *GameUsecase) StartGame(userID int) error {
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

func (u *GameUsecase) StopGame(userID int) error {
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

func (u *GameUsecase) SetDirection(userID int, direction int) error {
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

func (u *GameUsecase) SetSpeed(userID int, speed int) error {
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

func (u *GameUsecase) GetPlayer(userID int) (*model.Player, error) {
	res, err := u.game.GetPlayer(u.c, u.convertUserID(userID))
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	return u.convertPlayer(res.Player), nil
}

func (u *GameUsecase) GetPlayers(userID int) ([]*model.Player, error) {
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

func (u *GameUsecase) GetFood(userID int) ([]model.Food, error) {
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

func (u *GameUsecase) ListenEvents(userID int) (chan map[string]interface{}, error) {
	return make(chan map[string]interface{}), nil
}

func (u *GameUsecase) StopListenEvents(userID int) error {
	return nil
}

func (u *GameUsecase) convertUserID(id int) *game.UserID {
	return &game.UserID{UserId: int32(id)}
}

func (u *GameUsecase) convertPosition(pos *game.Position) model.Position {
	return model.Position{X: int(pos.X), Y: int(pos.Y)}
}

func (u *GameUsecase) convertPlayer(player *game.Player) *model.Player {
	return &model.Player{
		UserID:    int(player.UserId),
		Size:      int(player.Size),
		Direction: int(player.Direction),
		Speed:     int(player.Speed),
		Position:  u.convertPosition(player.Position),
	}
}

func (u *GameUsecase) convertFood(player *game.Food) model.Food {
	return model.Food{
		ID:       int(player.Id),
		Position: u.convertPosition(player.Position),
	}
}
