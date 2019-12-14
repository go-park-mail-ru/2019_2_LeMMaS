//+build wireinject
//go:generate wire

package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/delivery/grpc"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/usecase"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"os"
)

func NewAuthHandler() (*grpc.AuthHandler, error) {
	wire.Build(
		grpc.NewAuthHandler,
		usecase.NewAuthUsecase,
		repository.NewUserRepository,
		repository.NewSessionRepository,
		newRedis,
		NewLogger,
	)
	return &grpc.AuthHandler{}, nil
}

func NewLogger() (logger.Logger, error) {
	sentry, err := logger.NewSentryLogger()
	if err != nil {
		return nil, err
	}
	return logger.NewCombinedLogger(sentry, logger.NewStdoutLogger()), nil
}

func newRedis() (redis.Conn, error) {
	connection, err := redis.DialURL(os.Getenv("REDIS_DSN"))
	if err != nil {
		return nil, err
	}
	_, err = connection.Do("PING")
	if err != nil {
		return nil, err
	}
	return connection, nil
}
