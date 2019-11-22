package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Usecase interface {
	StartGame(userID int) error
	GameAlreadyStarted(userID int) bool
	StopGame(userID int) error

	SetDirection(userID int, direction int) error
	SetSpeed(userID int, speed int) error

	GetPlayers(userID int) map[int]*model.Player
	GetFood(userID int) map[int]model.Food
	GetEventsStream(userID int) chan model.GameEvent
}
