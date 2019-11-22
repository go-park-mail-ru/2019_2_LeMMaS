package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testUserID     = 1
	testRoomID     = 1
	testFoodAmount = 3
	testDirection  = 90
	testSpeed      = 100
)

func TestGameUsecase_StartGame(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	s.ExpectRepo().GetFoodInRange(testRoomID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()

	err := s.usecase.StartGame(testUserID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(s.usecase.GetPlayers(testUserID)))
	assert.Equal(t, testFoodAmount, len(s.usecase.GetFood(testUserID)))
}

func TestGameUsecase_SetDirectionAndSpeed(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	s.ExpectRepo().GetFoodInRange(testRoomID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()
	s.ExpectRepo().SetDirection(testRoomID, testUserID, testDirection).Return(nil)
	s.ExpectRepo().SetSpeed(testRoomID, testUserID, testSpeed).Return(nil)

	err := s.usecase.StartGame(testUserID)
	assert.NoError(t, err)
	assert.NoError(t, s.usecase.SetDirection(testUserID, testDirection))
	assert.NoError(t, s.usecase.SetSpeed(testUserID, testSpeed))
	player := s.usecase.GetPlayers(testUserID)[testUserID]
	assert.Equal(t, testDirection, player.Direction)
	assert.Equal(t, testSpeed, player.Speed)
}

func TestGameUsecase_PlayerMove(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	s.ExpectRepo().GetFoodInRange(testRoomID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()

	assert.NoError(t, s.usecase.StartGame(testUserID))
	initialPosition := s.usecase.GetPlayers(testUserID)[testUserID].Position
	events, err := s.usecase.ListenEvents(testUserID)
	assert.NoError(t, err)
	event := <-events
	if !assert.Equal(t, game.EventMove, event["type"]) {
		return
	}
	player := event["player"].(map[string]interface{})
	assert.Equal(t, testUserID, player["id"])
	assert.Greater(t, player["x"], initialPosition.X)
	assert.Equal(t, player["y"], initialPosition.Y)
	t.Logf("player moved (%v, %v) -> (%v, %v) in %v", initialPosition.X, initialPosition.Y, player["x"], player["y"], eventStreamRate)
}

func TestGameUsecase_EatFood(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	foodIDs := []int{3, 5}
	s.ExpectRepo().GetFoodInRange(testRoomID, gomock.Any(), gomock.Any()).Return(foodIDs, nil)
	s.ExpectRepo().DeleteFood(testRoomID, foodIDs).Return(nil)

	assert.NoError(t, s.usecase.StartGame(testUserID))
	events, err := s.usecase.ListenEvents(testUserID)
	assert.NoError(t, err)
	event := <-events
	if !assert.Equal(t, game.EventMove, event["type"]) {
		return
	}
	assert.Equal(t, foodIDs, event["eatenFood"])
}

type gameUsecaseTestSuite struct {
	t          *testing.T
	usecase    gameUsecase
	repository *game.MockRepository
}

func newGameUsecaseTestSuite(t *testing.T) *gameUsecaseTestSuite {
	s := gameUsecaseTestSuite{}
	s.t = t
	s.repository = game.NewMockRepository(gomock.NewController(t))
	s.usecase = gameUsecase{
		logger:           mock.NewMockLogger(),
		repository:       s.repository,
		roomsIDsByUserID: map[int]int{},
		eventsListeners:  map[int]map[int]chan model.GameEvent{},
	}
	s.initTestGame()
	return &s
}

func (s *gameUsecaseTestSuite) ExpectRepo() *game.MockRepositoryMockRecorder {
	return s.repository.EXPECT()
}

func (s *gameUsecaseTestSuite) initTestGame() {
	room := s.newTestRoom()
	s.ExpectRepo().GetAllRooms().Return([]*model.Room{})
	s.ExpectRepo().CreateRoom().Return(&room)
	s.ExpectRepo().AddPlayer(room.ID, gomock.Any()).Return(nil)
	s.ExpectRepo().AddFood(room.ID, gomock.Any()).Return(nil)
	s.ExpectRepo().DeleteFood(room.ID, gomock.Any()).Return(nil)
	s.ExpectRepo().GetRoomByID(room.ID).Return(&room).AnyTimes()
	s.ExpectRepo().SetPosition(room.ID, testUserID, gomock.Any()).Return(nil).AnyTimes()
}

func (s gameUsecaseTestSuite) newTestRoom() model.Room {
	player := model.Player{
		UserID:    testUserID,
		Direction: testDirection,
		Speed:     testSpeed,
		Position:  model.Position{X: game.MaxPositionX / 2, Y: game.MaxPositionY / 2},
	}
	food1 := model.Food{ID: 1, Position: model.Position{X: 10, Y: 10}}
	food2 := model.Food{ID: 2, Position: model.Position{X: 8, Y: 15}}
	food3 := model.Food{ID: 3, Position: model.Position{X: 20, Y: 50}}
	room := model.Room{
		ID: testRoomID,
		Players: map[int]*model.Player{
			testUserID: &player,
		},
		Food: map[int]model.Food{
			1: food1,
			2: food2,
			3: food3,
		},
	}
	return room
}
