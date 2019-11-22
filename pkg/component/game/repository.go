package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Repository interface {
	CreateRoom() *model.Room
	GetAllRooms() []*model.Room
	GetRoomByID(id int) *model.Room
	DeleteRoom(id int) error

	AddPlayer(roomID int, player model.Player) error
	DeletePlayer(roomID, userID int) error

	AddFood(roomID int, food []model.Food) error
	DeleteFood(roomID int, foodIDs []int) error
	GetFoodInRange(roomID int, topLeftPoint, bottomRightPoint model.Position) ([]int, error)

	SetDirection(roomID int, userID int, direction int) error
	SetSpeed(roomID int, userID int, speed int) error
	SetPosition(roomID int, userID int, position model.Position) error
}
