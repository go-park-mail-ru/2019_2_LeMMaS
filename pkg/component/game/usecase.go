package game

import "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"

type Usecase interface {
	SetDirection(direction int) error
	SetSpeed(speed int) error
	GetUpdatesStream() chan model.Position
}
