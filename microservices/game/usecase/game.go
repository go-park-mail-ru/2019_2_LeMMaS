package usecase

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/model"
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/proto"
	repo "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/repository"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"math"
	"math/rand"
	"time"
)

const (
	maxSpeed = 100
	minSpeed = 0

	maxDirection = 359
	minDirection = 0

	maxPlayersInRoom   = 5
	initialPlayerSize  = 40
	eatPlayerSizeKoeff = 5
	eatFoodSizeKoeff   = 2

	generatedFoodAmount = 20

	eventStreamRate = 50 * time.Millisecond
	speedKoeff      = float64(eventStreamRate/time.Millisecond) / 150
)

var (
	errGameNotStarted = errors.New("game not started")
)

type GameManager struct {
	repository       game.Repository
	roomsIDsByUserID map[int]int
	events           eventsDispatcher
}

func NewGameManager() *GameManager{
	repository := repo.NewRepository()

	return &GameManager{
		repository:       repository,
		roomsIDsByUserID: map[int]int{},
		events:           newEventsDispatcher(),
	}
}

func (g *GameManager) StartGame(ctx context.Context, userID *pb.UserID) (*pb.Error, error) {
	if g.getPlayerRoom(int(userID.UserID)) != nil {
		err := errors.New("start: game already started for this user")
		if err != nil {
			return &pb.Error{"start: game already started for this user"}, err
		}
		return &pb.Error{"ok"}, nil
	}

	room := g.getAvailableRoom()
	if room == nil {
		room = g.repository.CreateRoom()
		g.startEventsLoop(room.ID, ctx)
	}
	g.setPlayerRoom(int(userID.UserID), room.ID)

	player := g.newPlayer(int(userID.UserID))
	if err := g.repository.AddPlayer(room.ID, &player); err != nil {
		return &pb.Error{"cant't add player in room"}, err
	}

	food := g.generateFood()
	if err := g.repository.AddFood(room.ID, food); err != nil {
		return &pb.Error{"cant't add food in room"}, err
	}

	g.events.sendNewFood(room.ID, food)
	g.events.sendNewPlayer(room.ID, player)

	return &pb.Error{"ok"}, nil
}

func (g *GameManager) StopGame(ctx context.Context, userID *pb.UserID) (*pb.Error, error) {
	room := g.getPlayerRoom(int(userID.UserID))
	if room == nil {
		err := errors.New("stop: no game for this user to stop")
		return &pb.Error{"stop: no game for this user to stop"}, err
	}
	g.events.sendStop(room.ID, int(userID.UserID))
	g.StopListenEvents(ctx, userID)
	delete(g.roomsIDsByUserID, int(userID.UserID))
	if len(room.Players) == 1 {
		err := g.repository.DeleteRoom(room.ID)
		if err != nil {
			return &pb.Error{"error deleting room"}, err
		}
		return &pb.Error{"ok"}, nil
	}
	err := g.repository.DeletePlayer(room.ID, int(userID.UserID))
	if err != nil {
		return &pb.Error{"error deleting player"}, err
	}
	return &pb.Error{"ok"}, nil
}

func (g *GameManager) SetDirection(ctx context.Context, userANdDir *pb.UserAndDirection) (*pb.Error, error) {
	if userANdDir.Direction < minDirection || userANdDir.Direction > maxDirection {
		return &pb.Error{"error direction"}, fmt.Errorf("direction must be in range (%v, %v)", minDirection, maxDirection)
	}
	room := g.getPlayerRoom(int(userANdDir.UserID.UserID))
	if room == nil {
		return &pb.Error{"game not started"}, errGameNotStarted
	}
	err := g.repository.SetPlayerDirection(room.ID, int(userANdDir.UserID.UserID), int(userANdDir.Direction))
	if err != nil {
		return &pb.Error{"error don`t set player direcrion"}, err
	}
	return &pb.Error{"ok"}, nil
}

func (g *GameManager) SetSpeed(ctx context.Context, userAndSpeed *pb.UserAndSpeed) (*pb.Error, error) {
	if userAndSpeed.Speed < minSpeed || userAndSpeed.Speed > maxSpeed {
		return &pb.Error{"error speed"}, fmt.Errorf("speed must be in range (%v, %v)", minSpeed, maxSpeed)
	}
	room := g.getPlayerRoom(int(userAndSpeed.UserID.UserID))
	if room == nil {
		return &pb.Error{"error game not started"}, errGameNotStarted
	}
	err := g.repository.SetPlayerSpeed(room.ID, int(userAndSpeed.UserID.UserID), int(userAndSpeed.Speed))
	if err != nil {
		return &pb.Error{"error set player speed"}, err
	}
	return &pb.Error{"ok"}, nil
}

func (g *GameManager) GetPlayer(ctx context.Context, userID *pb.UserID) (*pb.Player, error) {
	players := g.getPlayerRoom(int(userID.UserID)).Players
	if player, ok := players[int(userID.UserID)]; ok {
		return player, nil
	}
	return &pb.Player{}, nil
}

func (g *GameManager) GetPlayers(ctx context.Context, userID *pb.UserID) (*pb.Players, error) {
	players := g.getPlayerRoom(int(userID.UserID)).Players
	result := make([]*model.Player, 0, len(players))
	for _, player := range players {
		result = append(result, player)
	}
	return result, nil
}

func (g *GameManager) GetFood(ctx context.Context, userID *pb.UserID) (*pb.Food, error) {
	foods := g.getPlayerRoom(int(userID.UserID)).Food
	result := make([]model.Food, 0, len(foods))
	for _, food := range foods {
		result = append(result, food)
	}
	return result, nil
}

func (g *GameManager) StopListenEvents(ctx context.Context, userID *pb.UserID) (*pb.Error, error) {
	room := g.getPlayerRoom(int(userID.UserID))
	if room == nil {
		return &pb.Error{"error game not started"}, errGameNotStarted
	}
	err := g.events.stopListen(room.ID, int(userID.UserID))
	if err != nil {
		return &pb.Error{"error stop listen"}, err
	}
	return &pb.Error{"ok"}, nil
}

func (g *GameManager) startEventsLoop(roomID int, ctx context.Context) {
	go func() {
		for range time.Tick(eventStreamRate) {
			room := g.repository.GetRoomByID(roomID)
			if room == nil {
				return
			}
			g.processPlayersMove(room, ctx)
		}
	}()
}

func (g *GameManager) processPlayersMove(room *model.Room, ctx context.Context) {
	for _, player := range room.Players {
		newPosition := g.getNewPosition(player)
		err := g.movePlayer(room, player, newPosition, ctx)
		if err != nil {
			log.Error(err)
		}
	}
}

func (g *GameManager) movePlayer(room *model.Room, player *model.Player, newPosition model.Position, ctx context.Context) error {
	if newPosition == player.Position {
		return nil
	}
	eatenPlayers, err := g.eatPlayers(room, player, newPosition, ctx)
	if err != nil {
		return err
	}
	eatenFood, err := g.eatFood(room, player, newPosition)
	if err != nil {
		return err
	}
	newSize := player.Size + len(eatenPlayers)*eatPlayerSizeKoeff + len(eatenFood)*eatFoodSizeKoeff
	if err := g.repository.SetPlayerSize(room.ID, player.UserID, newSize); err != nil {
		return err
	}
	if err := g.repository.SetPlayerPosition(room.ID, player.UserID, newPosition); err != nil {
		return err
	}
	g.events.sendMove(room.ID, player.UserID, newPosition, newSize, eatenFood)
	return nil
}

func (g *GameManager) newPlayer(userID int) model.Player {
	return model.Player{
		UserID: userID,
		Size:   initialPlayerSize,
		Position: model.Position{
			X: game.MaxPositionX / 2,
			Y: game.MaxPositionY / 2,
		},
	}
}

func (g *GameManager) getPlayerRoom(userID int) *model.Room {
	roomID, ok := g.roomsIDsByUserID[userID]
	if !ok {
		return nil
	}
	return g.repository.GetRoomByID(roomID)
}

func (g *GameManager) setPlayerRoom(userID, roomID int) {
	g.roomsIDsByUserID[userID] = roomID
}

func (g *GameManager) getAvailableRoom() *model.Room {
	availableRooms := g.repository.GetAllRooms()
	for _, room := range availableRooms {
		if len(room.Players) < maxPlayersInRoom {
			return room
		}
	}
	return nil
}

func (g *GameManager) getNewPosition(player *model.Player) model.Position {
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

func (g *GameManager) generateFood() []model.Food {
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

func (g *GameManager) eatFood(room *model.Room, player *model.Player, newPosition model.Position) ([]int, error) {
	eatenFood, err := g.getEatenFood(room.ID, player, newPosition)
	if err != nil {
		return nil, err
	}
	if err := g.repository.DeleteFood(room.ID, eatenFood); err != nil {
		return nil, err
	}
	return eatenFood, nil
}

func (g *GameManager) eatPlayers(room *model.Room, player *model.Player, newPosition model.Position, ctx context.Context) ([]int, error) {
	eatenPlayers, err := g.getEatenPlayers(room, player, newPosition)
	if err != nil {
		return nil, err
	}
	for _, eatenPlayer := range eatenPlayers {
		if _, err := g.StopGame(ctx, &pb.UserID{int32(eatenPlayer)}); err != nil {
			return nil, err
		}
	}
	return eatenPlayers, nil
}

func (g *GameManager) getEatenPlayers(room *model.Room, player *model.Player, position model.Position) ([]int, error) {
	p1, p2 := g.getEatingBound(player)
	playerIDs, err := g.repository.GetPlayersInRange(room.ID, p1, p2)
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

func (g *GameManager) getEatenFood(roomID int, player *model.Player, position model.Position) ([]int, error) {
	p1, p2 := g.getEatingBound(player)
	eatenFood, err := g.repository.GetFoodInRange(roomID, p1, p2)
	if err != nil {
		return nil, err
	}
	return eatenFood, nil
}

func (g *GameManager) getEatingBound(player *model.Player) (model.Position, model.Position) {
	r := player.Size/2 - 2
	pos := player.Position
	return model.Position{pos.X - r, pos.Y - r},
		model.Position{pos.X + r, pos.Y + r}
}
