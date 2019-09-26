package db

type UserID int32

type User struct {
	Id 				UserID `json:"-"`// pk
	Login 			string `json:"login"`// имя
	password		string `json:"-"` // пароль (не передается)
	AvatarAddress	string `json:"avatarAddress"`// адрес, где хранится аватар
	Role			string `json:"role"`// игрок залогинен или нет
}

type UserResults struct {
	Money	int `json:"money"`// количество монет
	Xp		int `json:"xp"`// количество очков
}

type TableScore struct {
	UserID	UserID `json:"-"`// id пользователя
	Login	string `json:"login"`// логин пользователя для отображения в таблице чемпионов
	Xp		int    `json:"xp"`// очки пользователя
}

type Shop struct {
	ItemID int  `json:"-"`// id товара
	Item string `json:"item"`// название товара
	Price int   `json:"price"`// цена товара
}