// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	http2 "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/access/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/access/usecase"
	http3 "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/user/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/user/repository"
	usecase2 "github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/user/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"os"
)

import (
	_ "github.com/jackc/pgx/stdlib"
)

// Injectors from wire.go:

func InitMiddleware() (http.Middleware, error) {
	echo := NewEcho()
	logger, err := NewLogger()
	if err != nil {
		return http.Middleware{}, err
	}
	middleware := http.NewMiddleware(echo, logger)
	return middleware, nil
}

func InitAccessHandler() (*http2.AccessHandler, error) {
	echo := NewEcho()
	csrfUsecase := usecase.NewCSRFUsecase()
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	accessHandler := http2.NewAccessHandler(echo, csrfUsecase, logger)
	return accessHandler, nil
}

func InitUserHandler() (*http3.UserHandler, error) {
	echo := NewEcho()
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewDatabaseRepository(db, logger)
	fileRepository := repository.NewFileRepository(logger)
	conn, err := NewRedis()
	if err != nil {
		return nil, err
	}
	sessionRepository := repository.NewSessionRepository(conn, logger)
	userUsecase := usecase2.NewUserUsecase(userRepository, fileRepository, sessionRepository)
	userHandler := http3.NewUserHandler(echo, userUsecase, logger)
	return userHandler, nil
}

// wire.go:

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
