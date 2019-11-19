package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Usecase interface {
	StartGame(user model.User) error
	SetDirection(user model.User, direction float64) error
	SetSpeed(user model.User, speed float64) error
	GetPlayers(user model.User) map[int]*model.Player
	GetFood(user model.User) map[int]*model.Position
	GetEventsStream(user model.User) chan model.GameEvent
}
