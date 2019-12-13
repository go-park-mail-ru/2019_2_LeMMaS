//+build wireinject
//go:generate wire

package auth

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/delivery/grpc"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/server"
	"github.com/google/wire"
)

func NewAuthHandler() (*grpc.AuthHandler, error) {
	wire.Build(
		grpc.NewAuthHandler,
		server.NewAuthServer,
		repository.NewUserRepository,
		repository.NewSessionRepository,
		NewRedis,
		logger.NewSentryLogger,
	)
	return &grpc.AuthHandler{}, nil
}

func NewRedis() (redis.Conn, error) {
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
