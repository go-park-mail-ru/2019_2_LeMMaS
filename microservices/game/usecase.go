package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/model"

type Usecase interface {
	StartGame(userID int) error
	StopGame(userID int) error

	SetDirection(userID int, direction int) error
	SetSpeed(userID int, speed int) error

	GetPlayer(userID int) *model.Player
	GetPlayers(userID int) []*model.Player
	GetFood(userID int) []model.Food

	ListenEvents(userID int) (chan map[string]interface{}, error)
	StopListenEvents(userID int) error
}
