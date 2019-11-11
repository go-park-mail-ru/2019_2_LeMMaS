package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/components/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/app/test/mock"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionRepository_AddSession(t *testing.T) {
	redis := redigomock.NewConn()
	repo := newTestSessionRepository(t, redis)
	userID := 3
	redis.Command(RedisCommandSet, test.SessionID, userID)
	err := repo.AddSession(test.SessionID, userID)
	assert.NoError(t, err)
	assert.NoError(t, redis.ExpectationsWereMet())
}

func newTestSessionRepository(t *testing.T, redis *redigomock.Conn) user.SessionRepository {
	logger := testMock.NewMockLogger(t)
	return NewSessionRepository(redis, logger)
}
