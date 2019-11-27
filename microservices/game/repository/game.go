package repository

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/quadtree"
)

type repository struct {
	rooms        map[int]*model.Room
	playersIndex map[int]*quadtree.Quadtree
	foodIndex    map[int]*quadtree.Quadtree
}

var errRoomNotFound = errors.New("room not found")

func NewRepository() game.Repository {
	return &repository{
		rooms:        map[int]*model.Room{},
		playersIndex: map[int]*quadtree.Quadtree{},
		foodIndex:    map[int]*quadtree.Quadtree{},
	}
}

type playerWrapper struct {
	player *model.Player
}

func (f playerWrapper) Point() orb.Point {
	return orb.Point{float64(f.player.Position.X), float64(f.player.Position.Y)}
}

type foodWrapper struct {
	food model.Food
}

func (f foodWrapper) Point() orb.Point {
	return orb.Point{float64(f.food.Position.X), float64(f.food.Position.Y)}
}

var roomIDCounter = 0

func (r *repository) CreateRoom() *model.Room {
	roomIDCounter++
	room := model.Room{
		ID:      roomIDCounter,
		Players: map[int]*model.Player{},
		Food:    map[int]model.Food{},
	}
	r.rooms[room.ID] = &room
	r.initRoomIndexes(room.ID)
	return &room
}

func (r *repository) initRoomIndexes(roomID int) {
	bound := orb.Bound{
		Min: orb.Point{0, 0},
		Max: orb.Point{game.MaxPositionX, game.MaxPositionY},
	}
	r.foodIndex[roomID] = quadtree.New(bound)
	r.playersIndex[roomID] = quadtree.New(bound)
}

func (r repository) GetRoomByID(id int) *model.Room {
	room, ok := r.rooms[id]
	if !ok {
		return nil
	}
	return room
}

func (r repository) GetAllRooms() []*model.Room {
	rooms := make([]*model.Room, 0, len(r.rooms))
	for _, room := range r.rooms {
		rooms = append(rooms, room)
	}
	return rooms
}

func (r *repository) DeleteRoom(id int) error {
	if !r.roomExists(id) {
		return errRoomNotFound
	}
	delete(r.rooms, id)
	return nil
}

func (r *repository) AddPlayer(roomID int, player *model.Player) error {
	if !r.roomExists(roomID) {
		return errRoomNotFound
	}
	r.rooms[roomID].Players[player.UserID] = player
	r.playersIndex[roomID].Add(playerWrapper{player})
	return nil
}

func (r *repository) DeletePlayer(roomID, userID int) error {
	if !r.roomExists(roomID) {
		return errRoomNotFound
	}
	player := r.rooms[roomID].Players[userID]
	r.playersIndex[roomID].Remove(playerWrapper{player}, nil)
	delete(r.rooms[roomID].Players, userID)
	return nil
}

func (r *repository) AddFood(roomID int, foods []model.Food) error {
	if !r.roomExists(roomID) {
		return errRoomNotFound
	}
	for _, food := range foods {
		r.rooms[roomID].Food[food.ID] = food
		r.foodIndex[roomID].Add(foodWrapper{food})
	}
	return nil
}

func (r *repository) DeleteFood(roomID int, foodIDs []int) error {
	if !r.roomExists(roomID) {
		return errRoomNotFound
	}
	for _, id := range foodIDs {
		food := r.rooms[roomID].Food[id]
		r.foodIndex[roomID].Remove(foodWrapper{food}, nil)
		delete(r.rooms[roomID].Food, id)
	}
	return nil
}

func (r *repository) GetPlayersInRange(roomID int, topLeftPoint, bottomRightPoint model.Position) ([]int, error) {
	if _, ok := r.playersIndex[roomID]; !ok {
		return nil, errRoomNotFound
	}
	wrappers := r.playersIndex[roomID].InBound(nil, r.newBound(topLeftPoint, bottomRightPoint))
	players := make([]int, 0, len(wrappers))
	for _, wrapper := range wrappers {
		players = append(players, wrapper.(playerWrapper).player.UserID)
	}
	return players, nil
}

func (r *repository) GetFoodInRange(roomID int, topLeftPoint, bottomRightPoint model.Position) ([]int, error) {
	if _, ok := r.foodIndex[roomID]; !ok {
		return nil, errRoomNotFound
	}
	wrappers := r.foodIndex[roomID].InBound(nil, r.newBound(topLeftPoint, bottomRightPoint))
	foods := make([]int, 0, len(wrappers))
	for _, wrapper := range wrappers {
		foods = append(foods, wrapper.(foodWrapper).food.ID)
	}
	return foods, nil
}

func (r repository) newBound(topLeftPoint, bottomRightPoint model.Position) orb.Bound {
	return orb.Bound{
		Min: orb.Point{float64(topLeftPoint.X), float64(topLeftPoint.Y)},
		Max: orb.Point{float64(bottomRightPoint.X), float64(bottomRightPoint.Y)},
	}
}

func (r *repository) SetPlayerDirection(roomID int, userID int, direction int) error {
	room := r.GetRoomByID(roomID)
	if room == nil {
		return errRoomNotFound
	}
	room.Players[userID].Direction = direction
	return nil
}

func (r *repository) SetPlayerSpeed(roomID int, userID int, speed int) error {
	room := r.GetRoomByID(roomID)
	if room == nil {
		return errRoomNotFound
	}
	room.Players[userID].Speed = speed
	return nil
}

func (r *repository) SetPlayerPosition(roomID int, userID int, position model.Position) error {
	room := r.GetRoomByID(roomID)
	if room == nil {
		return errRoomNotFound
	}
	room.Players[userID].Position = position
	return nil
}

func (r *repository) SetPlayerSize(roomID int, userID int, size int) error {
	room := r.GetRoomByID(roomID)
	if room == nil {
		return errRoomNotFound
	}
	room.Players[userID].Size = size
	return nil
}

func (r *repository) roomExists(roomID int) bool {
	_, exists := r.rooms[roomID]
	return exists
}
