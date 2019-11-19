package model

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

const (
	GameEventStart = "start"
	GameEventMove  = "move"
	GameEventError = "error"
)

type GameEvent = map[string]interface{}

type Player struct {
	Position  Position
	Direction int
	Speed     int
}
