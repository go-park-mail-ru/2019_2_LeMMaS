package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Repository interface {
	CreateRoom(userID int) *model.Room
	GetRoom(userID int) *model.Room
	DeleteRoom(userID int) error

	AddPlayer(room *model.Room, player model.Player)
	AddFood(room *model.Room, food []model.Food)
	GetFoodInRange(room *model.Room, topLeftPoint, bottomRightPoint model.Position) []model.Food

	SetDirection(userID int, direction int) error
	SetSpeed(userID int, speed int) error
	SetPosition(userID int, position model.Position) error
}
