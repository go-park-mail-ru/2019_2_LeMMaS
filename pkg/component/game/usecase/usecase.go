package usecase

import (
	"math/rand"
	"time"

	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

type gameUsecase struct {
	logger      logger.Logger
	playersByID map[int]model.User
}

func NewGameUsecase(logger logger.Logger) game.Usecase {
	return gameUsecase{logger: logger}
}

func (u gameUsecase) SetDirection(direction int) error {
	return nil
}

func (u gameUsecase) SetSpeed(speed int) error {
	return nil
}

func (u gameUsecase) GetUpdatesStream() chan model.Position {
	updates := make(chan model.Position)
	go func() {
		tick := time.Tick(300 * time.Millisecond)
		for range tick {
			updates <- model.Position{rand.Int(), rand.Int()}
		}
	}()
	return updates
}
