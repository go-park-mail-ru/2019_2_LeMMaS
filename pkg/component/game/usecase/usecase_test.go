package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userID = 1

func newTestGameUsecase() gameUsecase {
	return gameUsecase{
		repository:       repository.NewRepository(),
		roomsIDsByUserID: map[int]int{},
		gameStarted:      map[int]chan bool{},
	}
}

func TestGameUsecase_StartGame(t *testing.T) {
	u := newTestGameUsecase()
	err := u.StartGame(userID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(u.GetPlayers(userID)))
	assert.Equal(t, generatedFoodAmount, len(u.GetFood(userID)))
}

func TestGameUsecase_SetDirectionAndSpeed(t *testing.T) {
	u := newTestGameUsecase()
	err := u.StartGame(userID)
	assert.NoError(t, err)
	direction := 200
	speed := 100
	assert.NoError(t, u.SetDirection(userID, direction))
	assert.NoError(t, u.SetSpeed(userID, speed))
	player := u.GetPlayers(userID)[userID]
	assert.Equal(t, direction, player.Direction)
	assert.Equal(t, speed, player.Speed)
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
	assert.Less(t, position.X, game.MaxPositionX)
	assert.Less(t, position.Y, game.MaxPositionY)
}
