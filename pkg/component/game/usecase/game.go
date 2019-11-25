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
	maxSpeed = 100
	minSpeed = 0

	maxDirection = 359
	minDirection = 0

	initialPlayerSize = 40

	generatedFoodAmount = 10

	eventStreamRate = 1000 * time.Millisecond
	speedKoeff      = float64(eventStreamRate/time.Millisecond) / 150
)

var (
	errGameNotStarted = errors.New("game not started")
)

type gameUsecase struct {
	logger           logger.Logger
	repository       game.Repository
	roomsIDsByUserID map[int]int
	events           eventsDispatcher
}

func NewGameUsecase(repository game.Repository, logger logger.Logger) game.Usecase {
	return &gameUsecase{
		logger:           logger,
		repository:       repository,
		roomsIDsByUserID: map[int]int{},
		events:           newEventsDispatcher(),
	}
}

func (u *gameUsecase) StartGame(userID int) error {
	if u.getPlayerRoom(userID) != nil {
		return errors.New("start: game already started for this user")
	}

	room := u.getAvailableRoom()
	if room == nil {
		room = u.repository.CreateRoom()
		u.startEventsLoop(room)
	}
	u.setPlayerRoom(userID, room.ID)

	player := u.newPlayer(userID)
	if err := u.repository.AddPlayer(room.ID, player); err != nil {
		return err
	}

	food := u.generateFood()
	if err := u.repository.AddFood(room.ID, food); err != nil {
		return err
	}

	u.events.sendNewFood(room.ID, food)
	u.events.sendNewPlayer(room.ID, player)

	return nil
}

func (u *gameUsecase) StopGame(userID int) error {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errors.New("stop: no game for this user to stop")
	}
	u.StopListenEvents(userID)
	u.events.sendStop(room.ID, userID)
	delete(u.roomsIDsByUserID, userID)
	if len(room.Players) == 1 {
		return u.repository.DeleteRoom(room.ID)
	}
	err := u.repository.DeletePlayer(room.ID, userID)
	if err != nil {
		return err
	}
	return nil
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

func (u gameUsecase) GetPlayer(userID int) *model.Player {
	players := u.getPlayerRoom(userID).Players
	if player, ok := players[userID]; ok {
		return player
	}
	return nil
}

func (u gameUsecase) GetPlayers(userID int) []*model.Player {
	players := u.getPlayerRoom(userID).Players
	result := make([]*model.Player, 0, len(players))
	for _, player := range players {
		result = append(result, player)
	}
	return result
}

func (u gameUsecase) GetFood(userID int) []model.Food {
	foods := u.getPlayerRoom(userID).Food
	result := make([]model.Food, 0, len(foods))
	for _, food := range foods {
		result = append(result, food)
	}
	return result
}

func (u *gameUsecase) ListenEvents(userID int) (chan map[string]interface{}, error) {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return nil, errGameNotStarted
	}
	listener := u.events.Listen(room.ID, userID)
	return listener, nil
}

func (u *gameUsecase) StopListenEvents(userID int) error {
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errGameNotStarted
	}
	return u.events.stopListen(room.ID, userID)
}

func (u *gameUsecase) startEventsLoop(room *model.Room) {
	go func() {
		for range time.Tick(eventStreamRate) {
			err := u.processPlayersMove(room)
			if err != nil {
				return
			}
		}
	}()
}

func (u *gameUsecase) processPlayersMove(room *model.Room) error {
	for _, player := range room.Players {
		newPosition := u.getNextPlayerPosition(player)
		if newPosition != player.Position {
			err := u.repository.SetPosition(room.ID, player.UserID, newPosition)
			if err != nil {
				return err
			}
			eatenFoodIDs, err := u.eatFood(room.ID, player, newPosition)
			if err != nil {
				return err
			}
			u.events.sendMove(room.ID, player.UserID, newPosition, eatenFoodIDs)
		}
	}
	return nil
}

func (u *gameUsecase) newPlayer(userID int) model.Player {
	return model.Player{
		UserID: userID,
		Size:   initialPlayerSize,
		Position: model.Position{
			X: game.MaxPositionX / 2,
			Y: game.MaxPositionY / 2,
		},
	}
}

func (u *gameUsecase) getPlayerRoom(userID int) *model.Room {
	roomID, ok := u.roomsIDsByUserID[userID]
	if !ok {
		return nil
	}
	return u.repository.GetRoomByID(roomID)
}

func (u *gameUsecase) setPlayerRoom(userID, roomID int) {
	u.roomsIDsByUserID[userID] = roomID
}

func (u *gameUsecase) getAvailableRoom() *model.Room {
	availableRooms := u.repository.GetAllRooms()
	if len(availableRooms) == 0 {
		return nil
	}
	return availableRooms[0]
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

var foodIDCounter = 0

func (u gameUsecase) generateFood() []model.Food {
	foods := make([]model.Food, 0, generatedFoodAmount)
	for i := 0; i < generatedFoodAmount; i++ {
		foodIDCounter++
		foods = append(foods, model.Food{
			ID:       foodIDCounter,
			Position: model.Position{X: rand.Intn(game.MaxPositionX), Y: rand.Intn(game.MaxPositionY)},
		})
	}
	return foods
}

func (u *gameUsecase) eatFood(roomID int, player *model.Player, position model.Position) ([]int, error) {
	r := player.Size / 2
	eatenFoodIDs, err := u.repository.GetFoodInRange(
		roomID,
		model.Position{X: position.X - r, Y: position.Y - r},
		model.Position{X: position.X + r, Y: position.Y + r},
	)
	if err != nil {
		return nil, err
	}
	err = u.repository.DeleteFood(roomID, eatenFoodIDs)
	return eatenFoodIDs, err
}
