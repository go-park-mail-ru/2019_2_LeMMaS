package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/user"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
	"strconv"
)

const (
	redisCommandGet = "GET"
)

type sessionRepository struct {
	redis  redis.Conn
}

func NewSessionRepository(redis redis.Conn) user.SessionRepository {
	return &sessionRepository{
		redis,
	}
}

func (r *sessionRepository) GetUserBySession(sessionID string) (int, bool) {
	externalUserID, err := r.redis.Do(redisCommandGet, sessionID)
	if err != nil {
		log.Error(err)
		return 0, false
	}
	switch externalUserID.(type) {
	case []byte:
		userID, err := strconv.Atoi(string(externalUserID.([]byte)))
		if err != nil {
			log.Error(err)
			return 0, false
		}
		return userID, true
	}
	return 0, false
}
