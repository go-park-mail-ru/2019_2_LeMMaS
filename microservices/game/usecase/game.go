package usecase

import (
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/proto"
	repo "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"golang.org/x/net/context"
)

type GameManager struct {
	repository       game.Repository
	roomsIDsByUserID map[int]int
	events           eventsDispatcher
}

func NewGameManager() *GameManager{
	repository := repo.NewRepository()

	return &GameManager{
		repository:       repository,
		roomsIDsByUserID: map[int]int{},
		events:           newEventsDispatcher(),
	}
}

func (g GameManager) StartGame(context.Context, *pb.UserID) (*pb.Error, error) {
	panic("implement me")
}

func (g GameManager) StopGame(context.Context, *pb.UserID) (*pb.Error, error) {
	panic("implement me")
}

func (g GameManager) SetDirection(context.Context, *pb.UserAndDirection) (*pb.Error, error) {
	panic("implement me")
}

func (g GameManager) SetSpeed(context.Context, *pb.UserAndSpeed) (*pb.Error, error) {
	panic("implement me")
}

func (g GameManager) GetPlayer(context.Context, *pb.UserID) (*pb.Player, error) {
	panic("implement me")
}

func (g GameManager) GetPlayers(context.Context, *pb.UserID) (*pb.Players, error) {
	panic("implement me")
}

func (g GameManager) GetFood(context.Context, *pb.UserID) (*pb.Food, error) {
	panic("implement me")
}

func (g GameManager) StopListenEvents(context.Context, *pb.UserID) (*pb.Error, error) {
	panic("implement me")
}
