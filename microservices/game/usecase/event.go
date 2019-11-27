package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/model"
)

var errNoEventListeners = errors.New("no event listeners")

type eventsDispatcher struct {
	events map[int]roomEvents
}

func newEventsDispatcher() eventsDispatcher {
	return eventsDispatcher{
		events: map[int]roomEvents{},
	}
}

type roomEvents struct {
	listeners map[int]chan map[string]interface{}
}

func (u *eventsDispatcher) Listen(roomID, userID int) chan map[string]interface{} {
	listener := make(chan map[string]interface{})
	if _, ok := u.events[roomID]; !ok {
		u.events[roomID] = roomEvents{
			listeners: map[int]chan map[string]interface{}{},
		}
	}
	u.events[roomID].listeners[userID] = listener
	return listener
}

func (u *eventsDispatcher) stopListen(roomID, userID int) error {
	if _, ok := u.events[roomID]; !ok {
		return errNoEventListeners
	}
	if _, ok := u.events[roomID].listeners[userID]; !ok {
		return errNoEventListeners
	}
	delete(u.events[roomID].listeners, userID)
	return nil
}

func (u *eventsDispatcher) sendEvent(roomID int, event map[string]interface{}) {
	for _, listener := range u.events[roomID].listeners {
		listener <- event
	}
}

func (u *eventsDispatcher) sendStop(roomID, userID int) {
	u.sendEvent(roomID, map[string]interface{}{
		"type":    game.EventStop,
		"user_id": userID,
	})
}

func (u *eventsDispatcher) sendMove(roomID int, userID int, newPosition model.Position, newSize int, eatenFood []int) {
	u.sendEvent(roomID, map[string]interface{}{
		"type": game.EventMove,
		"player": map[string]interface{}{
			"id":   userID,
			"x":    newPosition.X,
			"y":    newPosition.Y,
			"size": newSize,
		},
		"eatenFood": eatenFood,
	})
}

func (u *eventsDispatcher) sendNewPlayer(roomID int, player model.Player) {
	u.sendEvent(roomID, map[string]interface{}{
		"type":   game.EventNewPlayer,
		"player": player,
	})
}

func (u *eventsDispatcher) sendNewFood(roomID int, food []model.Food) {
	u.sendEvent(roomID, map[string]interface{}{
		"type": game.EventNewFood,
		"food": food,
	})
}
