package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Usecase interface {
	SetDirection(user *model.User, direction int) error
	SetSpeed(user *model.User, speed int) error
	GetUpdatesStream(user *model.User) chan model.Position
}
