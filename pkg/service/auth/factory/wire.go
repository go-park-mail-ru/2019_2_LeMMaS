//+build wireinject
//go:generate wire

package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/delivery/grpc"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"os"
)

func NewAuthHandler() (*delivery.AuthHandler, error) {
	wire.Build(
		delivery.NewAuthHandler,
		usecase.NewAuthUsecase,
		repository.NewUserRepository,
		repository.NewSessionRepository,
		newUserClient,
		newRedis,
		NewLogger,
	)
	return &delivery.AuthHandler{}, nil
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

func newUserClient() (user.UserClient, error) {
	conn, err := newGRPC("user:" + os.Getenv("PORT"))
	return user.NewUserClient(conn), err
}

func newGRPC(url string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
