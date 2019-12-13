package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

const (
	set = "SET"
	get = "GET"
	del = "DEL"
)

type sessionRepo struct {
	redis  redis.Conn
	logger logger.Logger
}

func NewSessionRepository(redis redis.Conn, logger logger.Logger) auth.SessionRepository {
	return &sessionRepo{
		redis,
		logger,
	}
}

func (r *sessionRepo) Add(session string, user int) error {
	_, err := r.redis.Do(set, session, user)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}

func (r *sessionRepo) Get(session string) (int, bool) {
	value, err := r.redis.Do(get, session)
	if err != nil {
		r.logger.Error(err)
		return 0, false
	}
	switch value.(type) {
	case []byte:
		user, err := strconv.Atoi(string(value.([]byte)))
		if err != nil {
			r.logger.Error(err)
			return 0, false
		}
		return user, true
	}
	return 0, false
}

func (r *sessionRepo) Delete(session string) error {
	_, err := r.redis.Do(del, session)
	if err != nil {
		r.logger.Error(err)
	}
	return err
}
