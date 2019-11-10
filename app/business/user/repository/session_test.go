package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/test"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionRepository_AddSession(t *testing.T) {
	redis := redigomock.NewConn()
	repo := NewSessionRepository(redis)
	userID := 3
	redis.Command(RedisCommandSet, test.SessionID, userID)
	err := repo.AddSession(test.SessionID, userID)
	assert.NoError(t, err)
	assert.NoError(t, redis.ExpectationsWereMet())
}
