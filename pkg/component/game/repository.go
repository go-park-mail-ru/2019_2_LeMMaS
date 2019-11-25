//go:generate mockgen -source=$GOFILE -destination=repository_mock.go -package=$GOPACKAGE

package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Repository interface {
	CreateRoom() *model.Room
	GetAllRooms() []*model.Room
	GetRoomByID(id int) *model.Room
	DeleteRoom(id int) error

	AddPlayer(roomID int, player model.Player) error
	DeletePlayer(roomID, userID int) error
	SetPlayerDirection(roomID int, userID int, direction int) error
	SetPlayerSpeed(roomID int, userID int, speed int) error
	SetPlayerPosition(roomID int, userID int, position model.Position) error
	SetPlayerSize(roomID int, userID int, size int) error

	AddFood(roomID int, food []model.Food) error
	DeleteFood(roomID int, foodIDs []int) error
	GetFoodInRange(roomID int, topLeftPoint, bottomRightPoint model.Position) ([]int, error)
}
