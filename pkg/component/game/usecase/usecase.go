package usecase

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"math"
	"math/rand"
	"time"
)

const (
	maxSpeed = 100
	minSpeed = 0

	maxDirection = 359
	minDirection = 0

	playerSize = 10

	generatedFoodAmount = 10

	eventStreamRate = 1000 * time.Millisecond
)

type gameUsecase struct {
	repository  game.Repository
	gameStarted map[int]chan bool
}

func NewGameUsecase(roomRepository game.Repository) game.Usecase {
	return &gameUsecase{
		repository:  roomRepository,
		gameStarted: map[int]chan bool{},
	}
}

func (u *gameUsecase) StartGame(userID int) error {
	if u.repository.GetRoom(userID) != nil {
		return errors.New("start: game already started for this user")
	}
	room := u.repository.CreateRoom(userID)
	u.repository.AddPlayer(room, model.Player{
		UserID: userID,
		Position: model.Position{
			X: game.MaxPositionX / 2,
			Y: game.MaxPositionY / 2,
		},
	})
	u.repository.AddFood(room, u.generateFood())

	if gameStarted, ok := u.gameStarted[userID]; ok {
		gameStarted <- true
	}

	return nil
}

func (u *gameUsecase) StopGame(userID int) error {
	return u.repository.DeleteRoom(userID)
}

func (u *gameUsecase) SetDirection(userID int, direction int) error {
	if direction < minDirection || direction > maxDirection {
		return fmt.Errorf("direction must be in range (%v, %v)", minDirection, maxDirection)
	}
	return u.repository.SetDirection(userID, direction)
}

func (u *gameUsecase) SetSpeed(userID int, speed int) error {
	if speed < minSpeed || speed > maxSpeed {
		return fmt.Errorf("speed must be in range (%v, %v)", minSpeed, maxSpeed)
	}
	return u.repository.SetSpeed(userID, speed)
}

func (u gameUsecase) GetPlayers(userID int) map[int]*model.Player {
	return u.repository.GetRoom(userID).PlayersByID
}

func (u gameUsecase) GetFood(userID int) map[int]model.Food {
	return u.repository.GetRoom(userID).FoodByID
}

func (u *gameUsecase) GetEventsStream(userID int) chan model.GameEvent {
	events := make(chan model.GameEvent)
	go func() {
		if _, ok := u.gameStarted[userID]; !ok {
			u.gameStarted[userID] = make(chan bool)
		}
		<-u.gameStarted[userID]
		room := u.repository.GetRoom(userID)
		u.processEvents(room, userID, events)
	}()
	return events
}

func (u *gameUsecase) processEvents(room *model.Room, userID int, events chan model.GameEvent) {
	for range time.Tick(eventStreamRate) {
		player := room.PlayersByID[userID]
		newPosition := u.getNextPlayerPosition(player)
		if newPosition != player.Position {
			u.repository.SetPosition(userID, newPosition)

			eatenFood := u.getEatenFood(room, newPosition)
			eatenFoodIDs := make([]int, 0, len(eatenFood))
			for _, food := range eatenFood {
				eatenFoodIDs = append(eatenFoodIDs, food.ID)
			}

			events <- model.GameEvent{
				"type": game.EventMove,
				"players": map[string]interface{}{
					"id": userID,
					"x":  player.Position.X,
					"y":  player.Position.Y,
				},
				"eatenFood": eatenFoodIDs,
			}
		}
	}
}

func (u gameUsecase) getNextPlayerPosition(player *model.Player) model.Position {
	directionRadians := float64(player.Direction) * math.Pi / 180
	distance := float64(player.Speed) * float64(eventStreamRate/time.Millisecond) / 200
	deltaX := distance * math.Sin(directionRadians)
	deltaY := distance * math.Cos(directionRadians)
	oldPosition := player.Position
	newPosition := model.Position{
		X: int(math.Round(float64(oldPosition.X) + deltaX)),
		Y: int(math.Round(float64(oldPosition.Y) + deltaY)),
	}
	if newPosition.X > game.MaxPositionX {
		newPosition.X = game.MaxPositionX
	}
	if newPosition.Y > game.MaxPositionY {
		newPosition.Y = game.MaxPositionY
	}
	if newPosition.X < 0 {
		newPosition.X = 0
	}
	if newPosition.Y < 0 {
		newPosition.Y = 0
	}
	return newPosition
}

func (u gameUsecase) generateFood() []model.Food {
	foods := make([]model.Food, 0, generatedFoodAmount)
	for i := 0; i < generatedFoodAmount; i++ {
		foods = append(foods, model.Food{
			ID:       i + 1,
			Position: model.Position{X: rand.Intn(game.MaxPositionX), Y: rand.Intn(game.MaxPositionY)},
		})
	}
	return foods
}

func (u gameUsecase) getEatenFood(room *model.Room, position model.Position) []model.Food {
	return u.repository.GetFoodInRange(
		room,
		model.Position{X: position.X - playerSize/2, Y: position.Y - playerSize/2},
		model.Position{X: position.X + playerSize/2, Y: position.Y + playerSize/2},
	)
}
