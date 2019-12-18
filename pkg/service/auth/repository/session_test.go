package repository

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/consts"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test"
	testMock "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/test/mock"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

const (
	userID = 3
)

var errRedis = errors.New("redis error")

func newTestSessionRepository(t *testing.T, redis *redigomock.Conn) auth.SessionRepository {
	logger := testMock.NewMockLogger(t)
	return NewSessionRepository(redis, logger)
}

func TestSessionRepo_Add(t *testing.T) {
	redis := redigomock.NewConn()
	repo := newTestSessionRepository(t, redis)

	redis.Command(set, test.Session, userID)
	err := repo.Add(test.Session, userID)
	assert.NoError(t, err)

	redis.Command(set, test.Session, userID).ExpectError(errRedis)
	err = repo.Add(test.Session, userID)
	assert.Equal(t, consts.ErrStorageError, err)
}

func TestSessionRepo_Get(t *testing.T) {
	redis := redigomock.NewConn()
	repo := newTestSessionRepository(t, redis)

	redis.Command(get, test.Session).Expect([]byte(strconv.Itoa(userID)))
	id, ok := repo.Get(test.Session)
	if assert.True(t, ok) {
		assert.Equal(t, userID, id)
	}

	redis.Command(get, test.Session).ExpectError(errRedis)
	id, ok = repo.Get(test.Session)
	assert.False(t, ok)
}

func TestSessionRepo_Delete(t *testing.T) {
	redis := redigomock.NewConn()
	repo := newTestSessionRepository(t, redis)

	redis.Command(del, test.Session)
	err := repo.Delete(test.Session)
	assert.Nil(t, err)

	redis.Command(del, test.Session).ExpectError(errRedis)
	err = repo.Delete(test.Session)
	assert.Equal(t, consts.ErrStorageError, err)
}
