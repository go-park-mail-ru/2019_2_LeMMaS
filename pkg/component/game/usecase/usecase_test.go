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
)

func TestGameUsecase_StartGame(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	err := s.usecase.StartGame(testUserID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(s.usecase.GetPlayers(testUserID)))
	assert.Equal(t, testFoodAmount, len(s.usecase.GetFood(testUserID)))
}

func TestGameUsecase_SetDirectionAndSpeed(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	err := s.usecase.StartGame(testUserID)
	assert.NoError(t, err)

	direction := 90
	speed := 100
	s.ExpectRepo().SetDirection(testRoomID, testUserID, direction).Return(nil)
	s.ExpectRepo().SetSpeed(testRoomID, testUserID, speed).Return(nil)
	assert.NoError(t, s.usecase.SetDirection(testUserID, direction))
	assert.NoError(t, s.usecase.SetSpeed(testUserID, speed))
	player := s.usecase.GetPlayers(testUserID)[testUserID]
	assert.Equal(t, direction, player.Direction)
	assert.Equal(t, speed, player.Speed)
}

func TestGameUsecase_PlayerMove(t *testing.T) {
	s := newGameUsecaseTestSuite(t)
	direction := 90
	speed := 100
	s.ExpectRepo().SetDirection(testRoomID, testUserID, direction).Return(nil)
	s.ExpectRepo().SetSpeed(testRoomID, testUserID, speed).Return(nil)

	err := s.usecase.StartGame(testUserID)
	assert.NoError(t, err)
	initialPosition := s.usecase.GetPlayers(testUserID)[testUserID].Position
	assert.NoError(t, s.usecase.SetDirection(testUserID, direction))
	assert.NoError(t, s.usecase.SetSpeed(testUserID, speed))

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
	t.Logf("player moved from (%v, %v) to (%v, %v), time %v", initialPosition.X, initialPosition.Y, player["x"], player["y"], eventStreamRate)
}

//func TestGameUsecase_EatFood(t *testing.T) {
//	repo := game.NewMockRepository(gomock.NewController(t))
//	u := newTestGameUsecase(t, repo)
//
//	assert.NoError(t, u.StartGame(testUserID))
//	assert.NoError(t, u.SetDirection(testUserID, 90))
//	assert.NoError(t, u.SetSpeed(testUserID, 100))
//
//	events, err := u.ListenEvents(testUserID)
//	assert.NoError(t, err)
//	event := <-events
//	if !assert.Equal(t, game.EventMove, event["type"]) {
//		return
//	}
//	//event["eatenFood"]

type gameUsecaseTestSuite struct {
	t          *testing.T
	usecase    gameUsecase
	repository *game.MockRepository
}

func newGameUsecaseTestSuite(t *testing.T) gameUsecaseTestSuite {
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
	return s
}

func (s gameUsecaseTestSuite) ExpectRepo() *game.MockRepositoryMockRecorder {
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
	s.ExpectRepo().GetFoodInRange(room.ID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()
}

func (s gameUsecaseTestSuite) newTestRoom() model.Room {
	player := model.Player{
		UserID:    testUserID,
		Direction: 90,
		Speed:     100,
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
