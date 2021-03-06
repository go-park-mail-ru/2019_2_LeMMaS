package usecase

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
	"math"
	"math/rand"
	"time"
)

const (
	maxSpeed = 100
	minSpeed = 0

	maxDirection = 359
	minDirection = 0
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
		if err := u.StopGame(userID); err != nil {
			return err
		}
	}

	room := u.getAvailableRoom()
	if room == nil {
		room = u.repository.CreateRoom()
		u.startEventsLoop(room.ID)
	}
	u.setPlayerRoom(userID, room.ID)

	player := u.newPlayer(userID)
	if err := u.repository.AddPlayer(room.ID, &player); err != nil {
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
	u.events.sendStop(room.ID, userID)
	u.StopListenEvents(userID)
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
	return u.repository.SetPlayerDirection(room.ID, userID, direction)
}

func (u *gameUsecase) SetSpeed(userID int, speed int) error {
	if speed < minSpeed || speed > maxSpeed {
		return fmt.Errorf("speed must be in range (%v, %v)", minSpeed, maxSpeed)
	}
	room := u.getPlayerRoom(userID)
	if room == nil {
		return errGameNotStarted
	}
	return u.repository.SetPlayerSpeed(room.ID, userID, speed)
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

func (u *gameUsecase) startEventsLoop(roomID int) {
	go func() {
		for range time.Tick(game.EventStreamRate) {
			room := u.repository.GetRoomByID(roomID)
			if room == nil {
				return
			}
			u.processPlayersMove(room)
		}
	}()
}

func (u *gameUsecase) processPlayersMove(room *model.Room) {
	for _, player := range room.Players {
		newPosition := u.getNewPosition(player)
		err := u.movePlayer(room, player, newPosition)
		if err != nil {
			u.logger.Error(err)
		}
	}
}

func (u *gameUsecase) movePlayer(room *model.Room, player *model.Player, newPosition model.Position) error {
	if newPosition == player.Position {
		return nil
	}
	eatenPlayers, err := u.eatPlayers(room, player, newPosition)
	if err != nil {
		return err
	}
	eatenFood, err := u.eatFood(room, player, newPosition)
	if err != nil {
		return err
	}
	newSize := player.Size + len(eatenPlayers)*game.EatPlayerBonus + len(eatenFood)*game.EatFoodBonus
	if err := u.repository.SetPlayerSize(room.ID, player.UserID, newSize); err != nil {
		return err
	}
	if err := u.repository.SetPlayerPosition(room.ID, player.UserID, newPosition); err != nil {
		return err
	}
	u.events.sendMove(room.ID, player.UserID, newPosition, newSize, eatenFood)
	return nil
}

func (u *gameUsecase) newPlayer(userID int) model.Player {
	return model.Player{
		UserID: userID,
		Size:   game.InitialPlayerSize,
		Position: model.Position{
			X: game.FieldSizeX / 2,
			Y: game.FieldSizeY / 2,
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
	for _, room := range availableRooms {
		if len(room.Players) < game.MaxPlayersInRoom {
			return room
		}
	}
	return nil
}

func (u gameUsecase) getNewPosition(player *model.Player) model.Position {
	directionRadians := float64(player.Direction) * math.Pi / 180
	distance := float64(player.Speed) * float64(game.EventStreamRate/time.Millisecond) * game.Speed
	deltaX := distance * math.Sin(directionRadians)
	deltaY := -distance * math.Cos(directionRadians)
	oldPosition := player.Position
	newPosition := model.Position{
		X: int(math.Round(float64(oldPosition.X) + deltaX)),
		Y: int(math.Round(float64(oldPosition.Y) + deltaY)),
	}
	if newPosition.X > game.FieldSizeX {
		newPosition.X = game.FieldSizeX
	}
	if newPosition.Y > game.FieldSizeY {
		newPosition.Y = game.FieldSizeY
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
	foods := make([]model.Food, 0, game.FoodAmount)
	for i := 0; i < game.FoodAmount; i++ {
		foodIDCounter++
		foods = append(foods, model.Food{
			ID:       foodIDCounter,
			Position: model.Position{X: rand.Intn(game.FieldSizeX), Y: rand.Intn(game.FieldSizeY)},
		})
	}
	return foods
}

func (u *gameUsecase) eatFood(room *model.Room, player *model.Player, newPosition model.Position) ([]int, error) {
	eatenFood, err := u.getEatenFood(room.ID, player, newPosition)
	if err != nil {
		return nil, err
	}
	if err := u.repository.DeleteFood(room.ID, eatenFood); err != nil {
		return nil, err
	}
	return eatenFood, nil
}

func (u *gameUsecase) eatPlayers(room *model.Room, player *model.Player, newPosition model.Position) ([]int, error) {
	eatenPlayers, err := u.getEatenPlayers(room, player, newPosition)
	if err != nil {
		return nil, err
	}
	for _, eatenPlayer := range eatenPlayers {
		if err := u.StopGame(eatenPlayer); err != nil {
			return nil, err
		}
	}
	return eatenPlayers, nil
}

func (u *gameUsecase) getEatenPlayers(room *model.Room, player *model.Player, position model.Position) ([]int, error) {
	p1, p2 := u.getEatingBound(player)
	playerIDs, err := u.repository.GetPlayersInRange(room.ID, p1, p2)
	if err != nil {
		return nil, err
	}
	eatenPlayerIDs := playerIDs[:0]
	for _, playerID := range playerIDs {
		if playerID == player.UserID {
			continue
		}
		anotherPlayer, ok := room.Players[playerID]
		if !ok {
			return nil, errors.New("invalid eaten player id")
		}
		if player.Size > anotherPlayer.Size {
			eatenPlayerIDs = append(eatenPlayerIDs, playerID)
		}
	}
	return eatenPlayerIDs, nil
}

func (u *gameUsecase) getEatenFood(roomID int, player *model.Player, position model.Position) ([]int, error) {
	p1, p2 := u.getEatingBound(player)
	eatenFood, err := u.repository.GetFoodInRange(roomID, p1, p2)
	if err != nil {
		return nil, err
	}
	return eatenFood, nil
}

func (u *gameUsecase) getEatingBound(player *model.Player) (model.Position, model.Position) {
	r := player.Size/2 - 2
	pos := player.Position
	return model.Position{pos.X - r, pos.Y - r},
		model.Position{pos.X + r, pos.Y + r}
}
