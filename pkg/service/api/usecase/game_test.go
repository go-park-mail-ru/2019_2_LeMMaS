package usecase

//
//import (
//	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
//	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/game"
//	game2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
//	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//const (
//	userID     = 1
//	roomID     = 2
//	foodAmount = 3
//	direction  = 90
//	speed      = 100
//
//	testTimeout = 3 * time.Second
//)
//
//var s = gameUsecaseTestSuite{}
//
//func TestGameUsecase_StartGame(t *testing.T) {
//	s.StartTest(t)
//	s.ExpectRepo().GetFoodInRange(roomID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()
//
//	assert.NoError(t, s.server.StartGame(userID))
//	defer func() { assert.NoError(t, s.server.StopGame(userID)) }()
//
//	assert.Equal(t, 1, len(s.server.GetPlayers(userID)))
//	assert.Equal(t, foodAmount, len(s.server.GetFood(userID)))
//}
//
//func TestGameUsecase_SetDirectionAndSpeed(t *testing.T) {
//	s.StartTest(t)
//	s.ExpectRepo().GetFoodInRange(roomID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()
//	s.ExpectRepo().SetPlayerDirection(roomID, userID, direction).Return(nil)
//	s.ExpectRepo().SetPlayerSpeed(roomID, userID, speed).Return(nil)
//
//	assert.NoError(t, s.server.StartGame(userID))
//	defer func() { assert.NoError(t, s.server.StopGame(userID)) }()
//
//	assert.NoError(t, s.server.SetDirection(userID, direction))
//	assert.NoError(t, s.server.SetSpeed(userID, speed))
//	player := s.server.GetPlayer(userID)
//	assert.Equal(t, direction, player.Direction)
//	assert.Equal(t, speed, player.Speed)
//}
//
//func TestGameUsecase_PlayerMove(t *testing.T) {
//	timeout := time.After(testTimeout)
//	done := make(chan bool)
//	go func() {
//		s.StartTest(t)
//		s.ExpectRepo().GetFoodInRange(roomID, gomock.Any(), gomock.Any()).Return([]int{}, nil).AnyTimes()
//
//		assert.NoError(t, s.server.StartGame(userID))
//		defer func() { assert.NoError(t, s.server.StopGame(userID)) }()
//
//		initialPosition := s.server.GetPlayer(userID).Position
//		events, err := s.server.ListenEvents(userID)
//		assert.NoError(t, err)
//		event := <-events
//		if !assert.Equal(t, game2.EventMove, event["type"]) {
//			return
//		}
//		player := event["player"].(map[string]interface{})
//		assert.Equal(t, userID, player["id"])
//		assert.Greater(t, player["x"], initialPosition.X)
//		assert.Equal(t, player["y"], initialPosition.Y)
//		t.Logf("player moved (%v, %v) -> (%v, %v) in %v", initialPosition.X, initialPosition.Y, player["x"], player["y"], eventStreamRate)
//
//		done <- true
//	}()
//	select {
//	case <-timeout:
//		t.Fatal("test timed out, probably no game events were occured")
//	case <-done:
//	}
//}
//
//func TestGameUsecase_EatFood(t *testing.T) {
//	timeout := time.After(testTimeout)
//	done := make(chan bool)
//	go func() {
//		s.StartTest(t)
//		foodIDs := []int{3, 5}
//		s.ExpectRepo().GetFoodInRange(roomID, gomock.Any(), gomock.Any()).Return(foodIDs, nil).AnyTimes()
//		s.ExpectRepo().DeleteFood(roomID, gomock.Any()).Return(nil)
//
//		assert.NoError(t, s.server.StartGame(userID))
//		defer func() { assert.NoError(t, s.server.StopGame(userID)) }()
//
//		events, err := s.server.ListenEvents(userID)
//		assert.NoError(t, err)
//		event := <-events
//		if !assert.Equal(t, game2.EventMove, event["type"]) {
//			return
//		}
//		assert.Equal(t, foodIDs, event["eatenFood"])
//		done <- true
//	}()
//	select {
//	case <-timeout:
//		t.Fatal("test timed out, probably no game events were occured")
//	case <-done:
//	}
//}
//
//type gameUsecaseTestSuite struct {
//	t              *testing.T
//	server        game.Usecase
//	mockRepository *game2.MockRepository
//	controller     *gomock.Controller
//}
//
//func (s *gameUsecaseTestSuite) StartTest(t *testing.T) {
//	s.t = t
//	s.controller = gomock.NewController(t)
//	mockRepo := game2.NewMockRepository(s.controller)
//	s.server = NewGameUsecase(mockRepo, mock.NewMockLogger(t))
//	s.mockRepository = mockRepo
//	s.initTestGame()
//}
//
//func (s *gameUsecaseTestSuite) ExpectRepo() *game2.MockRepositoryMockRecorder {
//	return s.mockRepository.EXPECT()
//}
//
//func (s *gameUsecaseTestSuite) initTestGame() {
//	room := s.newTestRoom()
//	s.ExpectRepo().GetAllRooms().Return([]*model.Room{})
//	s.ExpectRepo().CreateRoom().Return(&room)
//	s.ExpectRepo().DeleteRoom(room.ID).Return(nil)
//	s.ExpectRepo().AddPlayer(room.ID, gomock.Any()).Return(nil)
//	s.ExpectRepo().GetPlayersInRange(room.ID, gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
//	s.ExpectRepo().AddFood(room.ID, gomock.Any()).Return(nil)
//	s.ExpectRepo().DeleteFood(room.ID, gomock.Any()).Return(nil).AnyTimes()
//	s.ExpectRepo().GetRoomByID(room.ID).Return(&room).AnyTimes()
//	s.ExpectRepo().SetPlayerPosition(room.ID, userID, gomock.Any()).Return(nil).AnyTimes()
//	s.ExpectRepo().SetPlayerSize(room.ID, userID, gomock.Any()).Return(nil).AnyTimes()
//}
//
//func (s gameUsecaseTestSuite) newTestRoom() model.Room {
//	player := model.Player{
//		UserID:    userID,
//		Direction: direction,
//		Speed:     speed,
//		Position:  model.Position{X: game2.MaxPositionX / 2, Y: game2.MaxPositionY / 2},
//	}
//	food1 := model.Food{ID: 1, Position: model.Position{X: 10, Y: 10}}
//	food2 := model.Food{ID: 2, Position: model.Position{X: 8, Y: 15}}
//	food3 := model.Food{ID: 3, Position: model.Position{X: 20, Y: 50}}
//	room := model.Room{
//		ID: roomID,
//		Players: map[int]*model.Player{
//			userID: &player,
//		},
//		Food: map[int]model.Food{
//			1: food1,
//			2: food2,
//			3: food3,
//		},
//	}
//	return room
//}
