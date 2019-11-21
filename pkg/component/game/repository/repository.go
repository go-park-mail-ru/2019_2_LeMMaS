package repository

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/quadtree"
)

type repository struct {
	roomsByID         map[int]*model.Room
	roomsIDsByUserID  map[int]int
	foodIndexByRoomID map[int]*quadtree.Quadtree
}

func NewRoomRepository() game.Repository {
	return &repository{
		roomsByID:         map[int]*model.Room{},
		roomsIDsByUserID:  map[int]int{},
		foodIndexByRoomID: map[int]*quadtree.Quadtree{},
	}
}

type foodWrapper struct {
	food model.Food
}

func (f foodWrapper) Point() orb.Point {
	return orb.Point{float64(f.food.Position.X), float64(f.food.Position.Y)}
}

var roomIDCounter = 0

func (r *repository) CreateRoom(userID int) *model.Room {
	roomIDCounter++
	room := model.Room{
		ID:          roomIDCounter,
		PlayersByID: map[int]*model.Player{},
		FoodByID:    map[int]model.Food{},
	}
	r.roomsByID[room.ID] = &room
	r.roomsIDsByUserID[userID] = room.ID
	r.foodIndexByRoomID[room.ID] = quadtree.New(orb.Bound{
		Min: orb.Point{0, 0},
		Max: orb.Point{model.MaxPositionX, model.MaxPositionY},
	})
	return &room
}

func (r repository) GetRoom(userID int) *model.Room {
	roomID, ok := r.roomsIDsByUserID[userID]
	if !ok {
		return nil
	}
	return r.roomsByID[roomID]
}

func (r *repository) DeleteRoom(userID int) error {
	roomID, ok := r.roomsIDsByUserID[userID]
	if !ok {
		return errors.New("stop: no game for this user to stop")
	}
	delete(r.roomsByID, roomID)
	delete(r.roomsIDsByUserID, userID)
	return nil
}

func (r *repository) AddPlayer(room *model.Room, player model.Player) {
	r.roomsByID[room.ID].PlayersByID[player.UserID] = &player
}

func (r *repository) AddFood(room *model.Room, foods []model.Food) {
	for _, food := range foods {
		r.roomsByID[room.ID].FoodByID[food.ID] = food
		r.foodIndexByRoomID[room.ID].Add(foodWrapper{food})
	}
}

func (r *repository) GetFoodInRange(room *model.Room, topLeftPoint, bottomRightPoint model.Position) []model.Food {
	wrappers := r.foodIndexByRoomID[room.ID].InBound(nil, orb.Bound{
		Min: orb.Point{float64(topLeftPoint.X), float64(topLeftPoint.Y)},
		Max: orb.Point{float64(bottomRightPoint.X), float64(bottomRightPoint.Y)},
	})
	foods := make([]model.Food, 0, len(wrappers))
	for _, wrapper := range wrappers {
		foods = append(foods, wrapper.(foodWrapper).food)
	}
	return foods
}

func (r *repository) SetDirection(userID int, direction int) error {
	room := r.GetRoom(userID)
	r.roomsByID[room.ID].PlayersByID[userID].Direction = direction
	return nil
}

func (r *repository) SetSpeed(userID int, speed int) error {
	room := r.GetRoom(userID)
	r.roomsByID[room.ID].PlayersByID[userID].Speed = speed
	return nil
}

func (r *repository) SetPosition(userID int, position model.Position) error {
	room := r.GetRoom(userID)
	r.roomsByID[room.ID].PlayersByID[userID].Position = position
	return nil
}
