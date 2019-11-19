package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

func TestGetNextPlayerPosition(t *testing.T) {
	u := gameUsecase{}
	initialPosition := model.Position{X: 0, Y: 0}
	player := model.Player{
		Position:  initialPosition,
		Direction: 90,
		Speed:     100,
	}
	position := u.GetNextPlayerPosition(player)
	t.Logf("got next position %v", position)
	assert.NotEqual(t, position, initialPosition)
	assert.Less(t, position.X, MaxPositionX)
	assert.Less(t, position.Y, MaxPositionY)
}
