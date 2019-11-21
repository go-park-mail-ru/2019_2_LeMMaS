package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Repository interface {
	CreateRoom(userID int) *model.Room
	GetAllRooms() []*model.Room
	GetRoomByID(id int) *model.Room
	DeleteRoom(userID int) error

	AddPlayer(room *model.Room, player model.Player)
	AddFood(room *model.Room, food []model.Food)
	GetFoodInRange(room *model.Room, topLeftPoint, bottomRightPoint model.Position) []model.Food

	SetDirection(room *model.Room, userID int, direction int) error
	SetSpeed(room *model.Room, userID int, speed int) error
	SetPosition(room *model.Room, userID int, position model.Position) error
}
