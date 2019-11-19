package model

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

const (
	GameEventStart = "start"
	GameEventMove  = "move"
	GameEventError = "error"
)

type GameEvent = map[string]interface{}

type Player struct {
	Position  Position
	Direction float64
	Speed     float64
}
