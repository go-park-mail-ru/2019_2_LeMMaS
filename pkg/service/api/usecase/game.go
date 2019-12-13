package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
)

type GameUsecase struct {
}

func NewGameUsecase() api.GameUsecase {
	return &GameUsecase{}
}

func (u *GameUsecase) StartGame(userID int) error {
	return nil

}
func (u *GameUsecase) StopGame(userID int) error {
	return nil
}

func (u *GameUsecase) SetDirection(userID int, direction int) error {
	return nil
}
func (u *GameUsecase) SetSpeed(userID int, speed int) error {
	return nil
}

func (u *GameUsecase) GetPlayer(userID int) *model.Player {
	return nil
}
func (u *GameUsecase) GetPlayers(userID int) []*model.Player {
	return nil
}
func (u *GameUsecase) GetFood(userID int) []model.Food {
	return nil
}

func (u *GameUsecase) ListenEvents(userID int) (chan map[string]interface{}, error) {
	return nil, nil
}
func (u *GameUsecase) StopListenEvents(userID int) error {
	return nil
}
