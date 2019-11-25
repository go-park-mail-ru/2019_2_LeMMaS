package model

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Player struct {
	UserID    int      `json:"user_id"`
	Size      int      `json:"size"`
	Position  Position `json:"position"`
	Direction int      `json:"-"`
	Speed     int      `json:"-"`
}

type Food struct {
	ID       int      `json:"id"`
	Position Position `json:"position"`
}

type Room struct {
	ID      int
	Players map[int]*Player
	Food    map[int]Food
}
