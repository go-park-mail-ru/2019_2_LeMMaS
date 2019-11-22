package usecase

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"math"
	"math/rand"
	"time"
)

const (
	maxSpeed   = 100
	minSpeed   = 0
	speedKoeff = float64(eventStreamRate / time.Millisecond / 500)

	maxDirection = 359
	minDirection = 0

	playerSize = 500

	generatedFoodAmount = 10

	eventStreamRate = 100 * time.Millisecond
)

var (
	errGameNotStarted = errors.New("game not started")
)

type gameUsecase struct {
	logger           logger.Logger
	repository       game.Repository
	roomsIDsByUserID map[int]int
	eventsListeners  map[int]map[int]chan model.GameEvent
}

func NewGameUsecase(repository game.Repository, logger logger.Logger) game.Usecase {
	return &gameUsecase{
		logger:           logger,
		repository:       repository,
		roomsIDsByUserID: map[int]int{},
		eventsListeners:  map[int]map[int]chan model.GameEvent{},
	}
}

func (u *gameUsecase) StartGame(userID int) error {
	if u.getPlayerRoom(userID) != nil {
		return errors.New("start: game already started for this user")
	}
	room := u.placePlayerInRoom(userID)
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
	if _, ok := u.eventsListeners[room.ID]; !ok {
		u.eventsListeners[room.ID] = map[int]chan model.GameEvent{}
	}
	u.eventsListeners[room.ID][userID] = listener
	return listener, nil
}

func (u *gameUsecase) StopListenEvents(userID int) error {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errGameNotStarted
	}
	if _, ok := u.eventsListeners[room.ID]; !ok {
		return errors.New("no event listeners")
	}
	delete(u.eventsListeners[room.ID], userID)
	return nil
}

func (u *gameUsecase) sendEvent(roomID int, event model.GameEvent) {
	for _, listener := range u.eventsListeners[roomID] {
		listener <- event
	}
}

func (u *gameUsecase) startRoomEventsLoop(room *model.Room) {
	go func() {
		for range time.Tick(eventStreamRate) {
			u.processPlayersMove(room)
		}
	}()
}

func (u *gameUsecase) processPlayersMove(room *model.Room) {
	for _, player := range room.Players {
		newPosition := u.getNextPlayerPosition(player)
		if newPosition != player.Position {
			err := u.repository.SetPosition(room.ID, player.UserID, newPosition)
			if err == repository.ErrRoomNotFound {
				break
			}
			eatenFoodIDs, err := u.eatFood(room.ID, newPosition)
			if err == repository.ErrRoomNotFound {
				break
			}
			event := model.GameEvent{
				"type": game.EventMove,
				"player": map[string]interface{}{
					"id": player.UserID,
					"x":  player.Position.X,
					"y":  player.Position.Y,
				},
				"eatenFood": eatenFoodIDs,
			}
			u.sendEvent(room.ID, event)
		}
	}
}

func (u *gameUsecase) getPlayerRoom(userID int) *model.Room {
	roomID, ok := u.roomsIDsByUserID[userID]
	if !ok {
		return nil
	}
	return u.repository.GetRoomByID(roomID)
}

func (u *gameUsecase) placePlayerInRoom(userID int) *model.Room {
	room := u.getPlayerRoom(userID)
	if room != nil {
		return room
	}
	availableRooms := u.repository.GetAllRooms()
	if len(availableRooms) == 0 {
		room = u.repository.CreateRoom()
		u.startRoomEventsLoop(room)
	} else {
		room = u.getAvailableRoom(availableRooms)
	}
	u.roomsIDsByUserID[userID] = room.ID
	return room
}

func (u gameUsecase) getAvailableRoom(rooms []*model.Room) *model.Room {
	return rooms[0]
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
