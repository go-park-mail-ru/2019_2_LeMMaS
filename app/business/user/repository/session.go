package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/logger"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

const (
	RedisCommandSet = "SET"
	RedisCommandGet = "GET"
	RedisCommandDel = "DEL"
)

type sessionRepository struct {
	redis redis.Conn
}

func NewSessionRepository(redis redis.Conn) *sessionRepository {
	return &sessionRepository{
		redis,
	}
}

func (r *sessionRepository) AddSession(sessionID string, userID int) error {
	_, err := r.redis.Do(RedisCommandSet, sessionID, userID)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *sessionRepository) GetUserBySession(sessionID string) (int, bool) {
	externalUserID, err := r.redis.Do(RedisCommandGet, sessionID)
	if err != nil {
		logger.Error(err)
		return 0, false
	}
	switch externalUserID.(type) {
	case []byte:
		userID, err := strconv.Atoi(string(externalUserID.([]byte)))
		if err != nil {
			logger.Error(err)
			return 0, false
		}
		return userID, true
	}
	return 0, false
}

func (r *sessionRepository) DeleteSession(sessionID string) error {
	_, err := r.redis.Do(RedisCommandDel, sessionID)
	if err != nil {
		logger.Error(err)
	}
	return err
}
