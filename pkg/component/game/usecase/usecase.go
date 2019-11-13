package usecase

import (
	"math/rand"
	"time"

	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

type gameUsecase struct {
	logger            logger.Logger
	positionsByUserID map[int]model.Position
}

func NewGameUsecase(logger logger.Logger) game.Usecase {
	return gameUsecase{
		logger:            logger,
		positionsByUserID: map[int]model.Position{},
	}
}

func (u gameUsecase) StartGame(user *model.User) error {
	u.positionsByUserID[user.ID] = model.Position{rand.Int(), rand.Int()}
	return nil
}

func (u gameUsecase) GetPlayerPosition(user *model.User) model.Position {
	return u.positionsByUserID[user.ID]
}

func (u gameUsecase) GetFoodsPositions(user *model.User) []model.Position {
	return []model.Position{{rand.Int(), rand.Int()}, {rand.Int(), rand.Int()}}
}

func (u gameUsecase) SetDirection(user *model.User, direction int) error {
	return nil
}

func (u gameUsecase) SetSpeed(user *model.User, speed int) error {
	return nil
}

func (u gameUsecase) GetUpdatesStream(user *model.User) chan model.Position {
	updates := make(chan model.Position)
	go func() {
		tick := time.Tick(4000 * time.Millisecond)
		for range tick {
			updates <- u.GetPlayerPosition(user)
		}
	}()
	return updates
}
