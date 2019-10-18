package model

type User struct {
	ID           int
	Email        string
	PasswordHash string
	Name         string
	AvatarPath   string
}
