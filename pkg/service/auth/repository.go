package auth

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

type SessionRepository interface {
	Add(session string, user int) error
	Get(session string) (int, bool)
	Delete(session string) error
}

type UserRepository interface {
	Create(email string, passwordHash string, name string) error
	GetByEmail(email string) (*model.User, error)
}
