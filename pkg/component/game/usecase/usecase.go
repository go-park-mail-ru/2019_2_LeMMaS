package usecase

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"math"
	"math/rand"
	"time"
)

const (
	maxSpeed   = 100
	minSpeed   = 0
	speedKoeff = float64(eventStreamRate/time.Millisecond) / 150

	maxDirection = 359
	minDirection = 0

	playerSize = 40

	generatedFoodAmount = 10

	eventStreamRate = 50 * time.Millisecond
)

var (
	errGameNotStarted = errors.New("game not started")
)

type gameUsecase struct {
	logger           logger.Logger
	repository       game.Repository
	roomsIDsByUserID map[int]int
	events           map[int]roomEvents
}

type roomEvents struct {
	listeners map[int]chan model.GameEvent
	stop      chan bool
}

func NewGameUsecase(repository game.Repository, logger logger.Logger) game.Usecase {
	return &gameUsecase{
		logger:           logger,
		repository:       repository,
		roomsIDsByUserID: map[int]int{},
		events:           map[int]roomEvents{},
	}
}

func (u *gameUsecase) StartGame(userID int) error {
	if u.getPlayerRoom(userID) != nil {
		return errors.New("start: game already started for this user")
	}
	room := u.getOrCreateRoom()
	u.roomsIDsByUserID[userID] = room.ID
	err := u.repository.AddPlayer(room.ID, model.Player{
		UserID: userID,
		Position: model.Position{
			X: game.MaxPositionX / 2,
			Y: game.MaxPositionY / 2,
		},
	})
	if err != nil {
		return err
	}
	return u.repository.AddFood(room.ID, u.generateFood())
}

func (u *gameUsecase) StopGame(userID int) error {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errors.New("stop: no game for this user to stop")
	}
	u.StopListenEvents(userID)
	delete(u.roomsIDsByUserID, userID)
	if len(room.Players) == 1 {
		return u.repository.DeleteRoom(room.ID)
	}
	return u.repository.DeletePlayer(room.ID, userID)
}

func (u *gameUsecase) SetDirection(userID int, direction int) error {
	if direction < minDirection || direction > maxDirection {
		return fmt.Errorf("direction must be in range (%v, %v)", minDirection, maxDirection)
	}
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errGameNotStarted
	}
	return u.repository.SetDirection(room.ID, userID, direction)
}

func (u *gameUsecase) SetSpeed(userID int, speed int) error {
	if speed < minSpeed || speed > maxSpeed {
		return fmt.Errorf("speed must be in range (%v, %v)", minSpeed, maxSpeed)
	}
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errGameNotStarted
	}
	return u.repository.SetSpeed(room.ID, userID, speed)
}

func (u gameUsecase) GetPlayers(userID int) map[int]*model.Player {
	return u.getPlayerRoom(userID).Players
}

func (u gameUsecase) GetFood(userID int) map[int]model.Food {
	return u.getPlayerRoom(userID).Food
}

func (u *gameUsecase) ListenEvents(userID int) (chan model.GameEvent, error) {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return nil, errGameNotStarted
	}
	listener := make(chan model.GameEvent)
	if _, ok := u.events[room.ID]; !ok {
		stop := make(chan bool)
		u.events[room.ID] = roomEvents{
			listeners: map[int]chan model.GameEvent{},
			stop:      stop,
		}
		u.startEventsLoop(room, stop)
	}
	u.events[room.ID].listeners[userID] = listener
	return listener, nil
}

func (u *gameUsecase) StopListenEvents(userID int) error {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errGameNotStarted
	}
	if _, ok := u.events[room.ID]; !ok {
		return errors.New("no event listeners")
	}
	if _, ok := u.events[room.ID].listeners[userID]; !ok {
		return errors.New("no event listeners")
	}
	delete(u.events[room.ID].listeners, userID)
	if len(u.events[room.ID].listeners) == 0 {
		u.stopEventsLoop(room)
	}
	return nil
}

func (u *gameUsecase) startEventsLoop(room *model.Room, stop chan bool) error {
	go func() {
		tick := time.Tick(eventStreamRate)
		for {
			select {
			case <-tick:
				err := u.processPlayersMove(room)
				if err != nil {
					return
				}
			case <-stop:
				return
			}
		}
	}()
	return nil
}

func (u *gameUsecase) stopEventsLoop(room *model.Room) {
	u.events[room.ID].stop <- true
}

func (u *gameUsecase) processPlayersMove(room *model.Room) error {
	for _, player := range room.Players {
		newPosition := u.getNextPlayerPosition(player)
		if newPosition != player.Position {
			err := u.repository.SetPosition(room.ID, player.UserID, newPosition)
			if err != nil {
				return err
			}
			eatenFoodIDs, err := u.eatFood(room.ID, newPosition)
			if err != nil {
				return err
			}
			event := model.GameEvent{
				"type": game.EventMove,
				"player": map[string]interface{}{
					"id": player.UserID,
					"x":  newPosition.X,
					"y":  newPosition.Y,
				},
				"eatenFood": eatenFoodIDs,
			}
			u.sendEvent(room.ID, event)
		}
	}
	return nil
}

func (u *gameUsecase) sendEvent(roomID int, event model.GameEvent) {
	for _, listener := range u.events[roomID].listeners {
		listener <- event
	}
}

func (u *gameUsecase) getPlayerRoom(userID int) *model.Room {
	roomID, ok := u.roomsIDsByUserID[userID]
	if !ok {
		return nil
	}
	return u.repository.GetRoomByID(roomID)
}

func (u *gameUsecase) getOrCreateRoom() *model.Room {
	var room *model.Room
	availableRooms := u.repository.GetAllRooms()
	if len(availableRooms) == 0 {
		room = u.repository.CreateRoom()
	} else {
		room = availableRooms[0]
	}
	return room
}

func (u gameUsecase) getNextPlayerPosition(player *model.Player) model.Position {
	directionRadians := float64(player.Direction) * math.Pi / 180
	distance := float64(player.Speed) * speedKoeff
	deltaX := distance * math.Sin(directionRadians)
	deltaY := -distance * math.Cos(directionRadians)
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

func (u *gameUsecase) eatFood(roomID int, playerPosition model.Position) ([]int, error) {
	eatenFoodIDs, err := u.repository.GetFoodInRange(
		roomID,
		model.Position{X: playerPosition.X - playerSize/2, Y: playerPosition.Y - playerSize/2},
		model.Position{X: playerPosition.X + playerSize/2, Y: playerPosition.Y + playerSize/2},
	)
	if err != nil {
		return nil, err
	}
	err = u.repository.DeleteFood(roomID, eatenFoodIDs)
	return eatenFoodIDs, err
}
