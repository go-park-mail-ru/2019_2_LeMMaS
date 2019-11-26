package auth

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/model"
)

type Repository interface {
	Create(email string, passwordHash string, name string) error
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}


type SessionRepository interface {
	AddSession(sessionID string, userID int) error
	GetUserBySession(sessionID string) (int, bool)
	DeleteSession(sessionID string) error
}

