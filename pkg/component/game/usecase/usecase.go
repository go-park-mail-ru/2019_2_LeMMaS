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
	MaxPositionX = 4000
	MaxPositionY = 2000

	MaxSpeed = 100
	MinSpeed = 0

	MaxDirection = 359
	MinDirection = 0

	EventStreamRate = 1000 * time.Millisecond
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
	u.playersByID[user.ID] = &model.Player{Position: model.Position{
		X: MaxPositionX / 2,
		Y: MaxPositionY / 2,
	}}
	u.foodByID = u.generateFood()
	u.gameStarted <- true
	return nil
}

func (u *gameUsecase) SetDirection(user model.User, direction int) error {
	if direction < MinDirection || direction > MaxDirection {
		return fmt.Errorf("direction must be in range (%v, %v)", MinDirection, MaxDirection)
	}
	u.playersByID[user.ID].Direction = direction
	return nil
}

func (u *gameUsecase) SetSpeed(user model.User, speed int) error {
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
			player := u.playersByID[user.ID]
			newPosition := u.GetNextPlayerPosition(*player)
			if newPosition != player.Position {
				player.Position = newPosition
				events <- model.GameEvent{
					"type": model.GameEventMove,
					"players": map[string]interface{}{
						"id": user.ID,
						"x":  player.Position.X,
						"y":  player.Position.Y,
					},
				}
			}
		}
	}()
	return events
}

func (u gameUsecase) GetNextPlayerPosition(player model.Player) model.Position {
	directionRadians := float64(player.Direction) * math.Pi / 180
	distance := float64(player.Speed) * float64(EventStreamRate/time.Millisecond) / 200
	deltaX := distance * math.Sin(directionRadians)
	deltaY := distance * math.Cos(directionRadians)
	oldPosition := player.Position
	newPosition := model.Position{
		X: int(math.Round(float64(oldPosition.X) + deltaX)),
		Y: int(math.Round(float64(oldPosition.Y) + deltaY)),
	}
	if newPosition.X > MaxPositionX {
		newPosition.X = MaxPositionX
	}
	if newPosition.Y > MaxPositionY {
		newPosition.Y = MaxPositionY
	}
	if newPosition.X < 0 {
		newPosition.X = 0
	}
	if newPosition.Y < 0 {
		newPosition.Y = 0
	}
	return newPosition
}

func (u gameUsecase) isValidPosition(position model.Position) bool {
	return position.X < MaxPositionX && position.Y < MaxPositionY
}

func (u gameUsecase) generateFood() map[int]*model.Position {
	food := map[int]*model.Position{}
	for i := 0; i < 10; i++ {
		food[i] = &model.Position{X: rand.Intn(MaxPositionX), Y: rand.Intn(MaxPositionY)}
	}
	return food
}
