package game

import "time"

const (
	FieldSizeX = 3000
	FieldSizeY = 3000

	Speed      = 1. / 700
	FoodAmount = 200

	EatFoodBonus   = 2
	EatPlayerBonus = 5

	EventStreamRate   = 50 * time.Millisecond
	MaxPlayersInRoom  = 5
	InitialPlayerSize = 20
)

const (
	EventStart     = "start"
	EventStop      = "stop"
	EventMove      = "move"
	EventNewPlayer = "new_player"
	EventNewFood   = "new_food"
	EventError     = "error"
)
