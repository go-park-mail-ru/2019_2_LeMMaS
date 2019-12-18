package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Food(t *testing.T) {
	r := NewRepository()

	room := r.CreateRoom()
	assert.NotNil(t, room)
	food := []model.Food{
		{ID: 1, Position: model.Position{X: 10, Y: 10}},
		{ID: 2, Position: model.Position{X: 8, Y: 15}},
		{ID: 3, Position: model.Position{X: 20, Y: 50}},
	}
	assert.NoError(t, r.AddFood(room.ID, food))

	foodIDs, err := r.GetFoodInRange(room.ID, model.Position{X: 5, Y: 5}, model.Position{X: 20, Y: 20})
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2}, foodIDs)

	assert.NoError(t, r.DeleteFood(room.ID, []int{1}))
	foodIDs, err = r.GetFoodInRange(room.ID, model.Position{X: 5, Y: 5}, model.Position{X: 20, Y: 20})
	assert.NoError(t, err)
	assert.Equal(t, []int{2}, foodIDs)
}
