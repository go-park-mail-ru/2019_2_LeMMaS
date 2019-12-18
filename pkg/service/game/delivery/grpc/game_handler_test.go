package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c context.Context

const (
	userID = 65
)

func TestGameHandler_StartGame(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	usecase.EXPECT().StartGame(userID).Return(nil)
	res, err := h.StartGame(c, h.convertUserID(userID))
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}

func TestGameHandler_StopGame(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	usecase.EXPECT().StopGame(userID).Return(nil)
	res, err := h.StopGame(c, h.convertUserID(userID))
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}

func TestGameHandler_GetPlayer(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	player := &model.Player{UserID: userID}
	usecase.EXPECT().GetPlayer(userID).Return(player)
	res, err := h.GetPlayer(c, h.convertUserID(userID))
	if assert.Nil(t, err) {
		assert.Equal(t, player.UserID, int(res.Player.UserId))
	}
}

func TestGameHandler_GetPlayers(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	expected := []*model.Player{{UserID: userID + 1}, {UserID: userID + 2}}
	usecase.EXPECT().GetPlayers(userID).Return(expected)
	res, err := h.GetPlayers(c, h.convertUserID(userID))
	if assert.Nil(t, err) {
		assert.Equal(t, len(expected), len(res.Players))
		for i, player := range res.Players {
			assert.Equal(t, expected[i].UserID, int(player.UserId))
		}
	}
}

func TestGameHandler_SetDirection(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	direction := 260
	usecase.EXPECT().SetDirection(userID, direction).Return(nil)
	res, err := h.SetDirection(c, &game.UserAndDirection{UserId: int32(userID), Direction: int32(direction)})
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}

func TestGameHandler_SetSpeed(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	speed := 90
	usecase.EXPECT().SetSpeed(userID, speed).Return(nil)
	res, err := h.SetSpeed(c, &game.UserAndSpeed{UserId: int32(userID), Speed: int32(speed)})
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}

func TestGameHandler_StopListenEvents(t *testing.T) {
	usecase := game.NewMockUsecase(gomock.NewController(t))
	h := NewGameHandler(usecase, mock.NewMockLogger(t))

	usecase.EXPECT().StopListenEvents(userID).Return(nil)
	res, err := h.StopListenEvents(c, h.convertUserID(userID))
	assert.Nil(t, err)
	assert.Empty(t, res.Error)
}
