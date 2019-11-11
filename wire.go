//+build wireinject

package main

import (
	accessHttp "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/access/delivery/http"
	accessUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/access/usecase"
	userHttp "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/user/delivery/http"
	userRepo "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/user/repository"
	userUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/user/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"os"
)

func InitMiddleware() (http.Middleware, error) {
	wire.Build(
		http.NewMiddleware,
		NewEcho,
		NewLogger,
	)
	return http.Middleware{}, nil
}

func InitAccessHandler() (*accessHttp.AccessHandler, error) {
	wire.Build(
		accessHttp.NewAccessHandler,
		accessUsecase.NewCSRFUsecase,
		NewEcho,
		NewLogger,
	)
	return &accessHttp.AccessHandler{}, nil
}

func InitUserHandler() (*userHttp.UserHandler, error) {
	wire.Build(
		userHttp.NewUserHandler,
		userUsecase.NewUserUsecase,
		userRepo.NewDatabaseRepository,
		userRepo.NewFileRepository,
		userRepo.NewSessionRepository,
		NewEcho,
		NewLogger,
		NewDB,
		NewRedis,
	)
	return &userHttp.UserHandler{}, nil
}

var echoInstance *echo.Echo = nil

func NewEcho() *echo.Echo {
	if echoInstance == nil {
		echoInstance = echo.New()
	}
	return echoInstance
}

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to postgres")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to postgres")
	}
	return db, nil
}

func NewRedis() (redis.Conn, error) {
	connection, err := redis.DialURL(os.Getenv("REDIS_DSN"))
	if err != nil {
		return connection, errors.Wrap(err, "cannot connect to redis")
	}
	_, err = connection.Do("PING")
	if err != nil {
		return connection, errors.Wrap(err, "cannot connect to redis")
	}
	return connection, nil
}

var loggerInstance *logger.Logger = nil

func NewLogger() (logger.Logger, error) {
	if loggerInstance == nil {
		echoLogger := logger.NewEchoLogger(NewEcho())
		sentryLogger, err := logger.NewSentryLogger()
		if err != nil {
			return nil, err
		}
		combinedLogger := logger.NewCombinedLogger(echoLogger, sentryLogger)
		loggerInstance = &combinedLogger
	}
	return *loggerInstance, nil
}
