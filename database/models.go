package db

import "sync"

var (
	PathAvatar = "usersAvatar/"
)

type AllUsers struct {
	Users []User
	mu    sync.RWMutex
}

type User struct {
	Id 				int    `json:"-"`// pk
	Login 			string `json:"login"`// имя
	Password		string `json:"-"` // пароль (не передается)
	Cookie 			string `json:"-"` // куки (не передаются)
	Email			string `json:"email"` // почта
	AvatarAddress	string `json:"avatarAddress"`// адрес, где хранится аватар
	Role			string `json:"role"`// игрок залогинен или нет
}

type UserResults struct {
	UserId  int `json:"-"`
	Money	int `json:"money"`// количество монет
	Xp		int `json:"xp"`// количество очков
}

type LeaderBoard struct {
	UserID	int    `json:"-"`// id пользователя
	Login	string `json:"login"`// логин пользователя для отображения в таблице чемпионов
	Xp		int    `json:"xp"`// очки пользователя
}

type Shop struct {
	ItemID int  `json:"-"`// id товара
	Item string `json:"item"`// название товара
	Price int   `json:"price"`// цена товара
}