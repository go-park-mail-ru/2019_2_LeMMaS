package model

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GameEvent = map[string]interface{}

type Player struct {
	UserID    int
	Position  Position
	Direction int
	Speed     int
}

type Food struct {
	ID       int
	Position Position
}

type Room struct {
	ID          int
	PlayersByID map[int]*Player
	FoodByID    map[int]Food
}
