//go:generate mockgen -source=$GOFILE -destination=usecase_mock.go -package=$GOPACKAGE

package api

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"io"
)

type AuthUsecase interface {
	Login(email, password string) (sessionID string, err error)
	Logout(sessionID string) error
	Register(email, password, name string) error
	GetUserBySession(sessionID string) (*model.User, error)
}

type UserUsecase interface {
	GetAll() ([]*model.User, error)
	GetByID(id int) (*model.User, error)
	Update(id int, password, name string) error
	UpdateAvatar(id int, avatar io.Reader) error
	GetSpecialAvatar(name string) (string, error)
}

type CsrfUsecase interface {
	CreateTokenBySession(sessionID string) (string, error)
	CheckTokenBySession(token string, sessionID string) (bool, error)
}

type GameUsecase interface {
	StartGame(userID int) error
	StopGame(userID int) error

	SetDirection(userID int, direction int) error
	SetSpeed(userID int, speed int) error

	GetPlayer(userID int) *model.Player
	GetPlayers(userID int) []*model.Player
	GetFood(userID int) []model.Food

	ListenEvents(userID int) (chan map[string]interface{}, error)
	StopListenEvents(userID int) error
}
