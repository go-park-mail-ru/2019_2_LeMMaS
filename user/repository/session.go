package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	"github.com/gomodule/redigo/redis"
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
	_, err := r.redis.Do("SET", sessionID, userID)
	if err != nil {
		logger.Error(err)
	}
	return err
}

func (r *sessionRepository) GetUserBySession(sessionID string) (int, bool) {
	userID, err := r.redis.Do("GET", sessionID)
	if err != nil {
		logger.Error(err)
		return 0, false
	}
	return userID.(int), true
}

func (r *sessionRepository) DeleteSession(sessionID string) error {
	_, err := r.redis.Do("DELETE", sessionID)
	if err != nil {
		logger.Error(err)
	}
	return err
}
