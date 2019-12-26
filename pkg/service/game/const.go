package game

import "time"

const (
	FieldSizeX = 1050
	FieldSizeY = 1050

	Speed      = 1. / 700
	FoodAmount = 80

	EatFoodBonus   = 2
	EatPlayerBonus = 5

	EventStreamRate   = 50 * time.Millisecond
	MaxPlayersInRoom  = 7
	InitialPlayerSize = 20
)

const (
	EventStart     = "start"
	EventStop      = "stop"
	EventError     = "error"
	EventMove      = "move"
	EventNewPlayer = "new_player"
	EventNewFood   = "new_food"
)
