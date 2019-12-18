package model

type User struct {
	ID           int    `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"-" db:"password_hash"`
	Name         string `json:"name" db:"name"`
	AvatarPath   string `json:"avatar_path" db:"avatar_path"`
}
