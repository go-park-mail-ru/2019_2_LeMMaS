package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Usecase interface {
	StartGame(user *model.User) error
	GetPlayerPosition(user *model.User) model.Position
	GetFoodsPositions(user *model.User) []model.Position
	SetDirection(user *model.User, direction int) error
	SetSpeed(user *model.User, speed int) error
	GetUpdatesStream(user *model.User) chan model.Position
}
