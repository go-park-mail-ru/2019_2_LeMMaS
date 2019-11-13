package usecase

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

const EventStreamRate = 2000 * time.Millisecond

const MaxSpeed = 100
const MinSpeed = 0

type gameUsecase struct {
	logger      logger.Logger
	infoByUser  map[int]*userInfo
	gameStarted chan bool
}

type userInfo struct {
	Position  model.Position
	Food      []model.Position
	Direction float64
	Speed     float64
}

func NewGameUsecase(logger logger.Logger) game.Usecase {
	return &gameUsecase{
		logger:      logger,
		infoByUser:  map[int]*userInfo{},
		gameStarted: make(chan bool),
	}
}

func (u gameUsecase) StartGame(user model.User) error {
	u.infoByUser[user.ID] = &userInfo{
		Position:  model.Position{X: 0, Y: 0},
		Food:      []model.Position{{rand.Float64(), rand.Float64()}, {rand.Float64(), rand.Float64()}},
		Direction: 0,
		Speed:     0,
	}
	u.gameStarted <- true
	return nil
}

func (u *gameUsecase) SetDirection(user model.User, direction float64) error {
	if direction < 0 || direction > 359 {
		return errors.New("directions must be in range (0, 359)")
	}
	u.infoByUser[user.ID].Direction = direction
	return nil
}

func (u *gameUsecase) SetSpeed(user model.User, speed float64) error {
	if speed < MinSpeed || speed > MaxSpeed {
		return fmt.Errorf("speed must be in range (%v, %v)", MinSpeed, MaxSpeed)
	}
	u.infoByUser[user.ID].Speed = speed
	return nil
}

func (u gameUsecase) GetPlayerPosition(user model.User) model.Position {
	return u.infoByUser[user.ID].Position
}

func (u gameUsecase) GetFoodsPositions(user model.User) []model.Position {
	return u.infoByUser[user.ID].Food
}

func (u *gameUsecase) GetEventStream(user model.User) chan game.Event {
	updates := make(chan game.Event)
	go func() {
		<-u.gameStarted
		tick := time.Tick(EventStreamRate)
		for range tick {
			u.updatePlayerPosition(user)
			info := u.infoByUser[user.ID]
			updates <- game.Event{
				Type: game.EventTypeMove,
				Body: map[string]interface{}{
					"x": info.Position.X,
					"y": info.Position.Y,
				},
			}
		}
	}()
	return updates
}

func (u *gameUsecase) updatePlayerPosition(user model.User) {
	info := u.infoByUser[user.ID]
	directionRadians := info.Direction * math.Pi / 180
	distance := info.Speed * float64(EventStreamRate/time.Millisecond)
	deltaX := distance * math.Sin(directionRadians)
	deltaY := distance * math.Cos(directionRadians)
	oldPosition := info.Position
	newPosition := model.Position{
		X: oldPosition.X + deltaX,
		Y: oldPosition.Y + deltaY,
	}
	newPosition.X = math.Round(newPosition.X*100) / 100
	newPosition.Y = math.Round(newPosition.Y*100) / 100
	info.Position = newPosition
}
