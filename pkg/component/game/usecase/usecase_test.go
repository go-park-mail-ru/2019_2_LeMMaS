package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestGameUsecase() gameUsecase {
	return gameUsecase{
		repository:  repository.NewRoomRepository(),
		gameStarted: map[int]chan bool{},
	}
}

func TestGameUsecase_StartGame(t *testing.T) {
	u := newTestGameUsecase()
	userID := 1
	err := u.StartGame(userID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(u.GetPlayers(userID)))
	assert.Equal(t, generatedFoodAmount, len(u.GetFood(userID)))
}

func TestGetNextPlayerPosition(t *testing.T) {
	u := newTestGameUsecase()
	initialPosition := model.Position{X: 0, Y: 0}
	player := model.Player{
		Position:  initialPosition,
		Direction: 90,
		Speed:     100,
	}
	position := u.getNextPlayerPosition(&player)
	t.Logf("got next position %v", position)
	assert.NotEqual(t, position, initialPosition)
	assert.Less(t, position.X, model.MaxPositionX)
	assert.Less(t, position.Y, model.MaxPositionY)
}
