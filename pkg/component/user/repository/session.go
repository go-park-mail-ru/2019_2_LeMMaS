package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

const (
	redisCommandSet = "SET"
	redisCommandGet = "GET"
	redisCommandDel = "DEL"
)

type sessionRepository struct {
	redis  redis.Conn
	logger logger.Logger
}

func NewSessionRepository(redis redis.Conn, logger logger.Logger) user.SessionRepository {
	return &sessionRepository{
		redis,
		logger,
	}
}

func (r *sessionRepository) AddSession(sessionID string, userID int) error {
	_, err := r.redis.Do(redisCommandSet, sessionID, userID)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}

func (r *sessionRepository) GetUserBySession(sessionID string) (int, bool) {
	externalUserID, err := r.redis.Do(redisCommandGet, sessionID)
	if err != nil {
		r.logger.Error(err)
		return 0, false
	}
	switch externalUserID.(type) {
	case []byte:
		userID, err := strconv.Atoi(string(externalUserID.([]byte)))
		if err != nil {
			r.logger.Error(err)
			return 0, false
		}
		return userID, true
	}
	return 0, false
}

func (r *sessionRepository) DeleteSession(sessionID string) error {
	_, err := r.redis.Do(redisCommandDel, sessionID)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}
