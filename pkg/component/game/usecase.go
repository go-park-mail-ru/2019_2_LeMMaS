package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

const EventTypeMove = "move"

type Event struct {
	Type string                 `json:"type"`
	Body map[string]interface{} `json:"body"`
}

type Usecase interface {
	StartGame(user model.User) error
	SetDirection(user model.User, direction float64) error
	SetSpeed(user model.User, speed float64) error
	GetPlayerPosition(user model.User) model.Position
	GetFoodsPositions(user model.User) []model.Position
	GetEventsStream(user model.User) chan Event
}
