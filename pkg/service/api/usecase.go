//go:generate mockgen -source=$GOFILE -destination=usecase_mock.go -package=$GOPACKAGE

package api

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"io"
)

type AuthUsecase interface {
	Login(email, password string) (session string, err error)
	Logout(session string) error
	Register(email, password, name string) error
	GetUserID(session string) (int, error)
	GetPasswordHash(password string) (string, error)
}

type UserUsecase interface {
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	Update(id int, passwordHash, name string) error
	UpdateAvatar(id int, avatar io.Reader) error
	GetSpecialAvatar(name string) (string, error)
}

type CsrfUsecase interface {
	CreateTokenBySession(session string) (string, error)
	CheckTokenBySession(token string, session string) (bool, error)
}

type GameUsecase interface {
	StartGame(userID int) error
	StopGame(userID int) error

	SetDirection(userID int, direction int) error
	SetSpeed(userID int, speed int) error

	GetPlayer(userID int) (*model.Player, error)
	GetPlayers(userID int) ([]*model.Player, error)
	GetFood(userID int) ([]model.Food, error)

	ListenEvents(userID int) (chan map[string]interface{}, error)
	StopListenEvents(userID int) error
}
