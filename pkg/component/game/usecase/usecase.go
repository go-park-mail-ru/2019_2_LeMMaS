package usecase

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

const (
	MaxSpeed = 100
	MinSpeed = 0

	MaxDirection = 359
	MinDirection = 0

	EventStreamRate = 2000 * time.Millisecond
)

type gameUsecase struct {
	logger logger.Logger

	playersByID map[int]*model.Player
	foodByID    map[int]*model.Position
	gameStarted chan bool
}

func NewGameUsecase(logger logger.Logger) game.Usecase {
	return &gameUsecase{
		logger:      logger,
		playersByID: map[int]*model.Player{},
		foodByID:    map[int]*model.Position{},
		gameStarted: make(chan bool),
	}
}

func (u *gameUsecase) StartGame(user model.User) error {
	u.playersByID[user.ID] = &model.Player{}
	u.foodByID = u.generateFood()
	u.gameStarted <- true
	return nil
}

func (u *gameUsecase) SetDirection(user model.User, direction float64) error {
	if direction < MinDirection || direction > MaxDirection {
		return fmt.Errorf("direction must be in range (%v, %v)", MinDirection, MaxDirection)
	}
	u.playersByID[user.ID].Direction = direction
	return nil
}

func (u *gameUsecase) SetSpeed(user model.User, speed float64) error {
	if speed < MinSpeed || speed > MaxSpeed {
		return fmt.Errorf("speed must be in range (%v, %v)", MinSpeed, MaxSpeed)
	}
	u.playersByID[user.ID].Speed = speed
	return nil
}

func (u gameUsecase) GetPlayers(user model.User) map[int]*model.Player {
	return u.playersByID
}

func (u gameUsecase) GetFood(user model.User) map[int]*model.Position {
	return u.foodByID
}

func (u *gameUsecase) GetEventsStream(user model.User) chan model.GameEvent {
	events := make(chan model.GameEvent)
	go func() {
		<-u.gameStarted
		tick := time.Tick(EventStreamRate)
		for range tick {
			u.updatePlayerPosition(user)
			player := u.playersByID[user.ID]
			events <- model.GameEvent{
				"type": model.GameEventMove,
				"players": map[string]interface{}{
					"id": user.ID,
					"x":  player.Position.X,
					"y":  player.Position.Y,
				},
			}
		}
	}()
	return events
}

func (u gameUsecase) generateFood() map[int]*model.Position {
	food := map[int]*model.Position{}
	for i := 0; i < 10; i++ {
		food[i] = &model.Position{X: rand.Float64(), Y: rand.Float64()}
	}
	return food
}

func (u *gameUsecase) updatePlayerPosition(user model.User) {
	player := u.playersByID[user.ID]
	directionRadians := player.Direction * math.Pi / 180
	distance := player.Speed * float64(EventStreamRate/time.Millisecond)
	deltaX := distance * math.Sin(directionRadians)
	deltaY := distance * math.Cos(directionRadians)
	oldPosition := player.Position
	newPosition := model.Position{
		X: oldPosition.X + deltaX,
		Y: oldPosition.Y + deltaY,
	}
	newPosition.X = math.Round(newPosition.X*100) / 100
	newPosition.Y = math.Round(newPosition.Y*100) / 100
	player.Position = newPosition
}
