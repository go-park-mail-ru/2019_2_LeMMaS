// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package bootstrap

import (
	http2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/access/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/access/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game/delivery/ws"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game/repository"
	usecase2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/game/usecase"
	http3 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user/delivery/http"
	repository2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user/repository"
	usecase3 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"os"
)

// Injectors from wire.go:

func NewMiddleware() (http.CommonMiddlewaresHandler, error) {
	echo := NewEcho()
	logger, err := NewLogger()
	if err != nil {
		return http.CommonMiddlewaresHandler{}, err
	}
	commonMiddlewaresHandler := http.NewCommonMiddlewaresHandler(echo, logger)
	return commonMiddlewaresHandler, nil
}

func NewAccessHandler() (*http2.AccessHandler, error) {
	echo := NewEcho()
	csrfUsecase := usecase.NewCSRFUsecase()
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	accessHandler := http2.NewAccessHandler(echo, csrfUsecase, logger)
	return accessHandler, nil
}

func NewGameHandler() (*ws.GameHandler, error) {
	echo := NewEcho()
	roomRepository := repository.NewRoomRepository()
	gameUsecase := usecase2.NewGameUsecase(roomRepository)
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	userRepository := repository2.NewDatabaseRepository(db, logger)
	fileRepository := repository2.NewS3Repository(logger)
	conn, err := NewRedis()
	if err != nil {
		return nil, err
	}
	sessionRepository := repository2.NewSessionRepository(conn, logger)
	userUsecase := usecase3.NewUserUsecase(userRepository, fileRepository, sessionRepository)
	gameHandler := ws.NewGameHandler(echo, gameUsecase, userUsecase, logger)
	return gameHandler, nil
}

func NewUserHandler() (*http3.UserHandler, error) {
	echo := NewEcho()
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	userRepository := repository2.NewDatabaseRepository(db, logger)
	fileRepository := repository2.NewS3Repository(logger)
	conn, err := NewRedis()
	if err != nil {
		return nil, err
	}
	sessionRepository := repository2.NewSessionRepository(conn, logger)
	userUsecase := usecase3.NewUserUsecase(userRepository, fileRepository, sessionRepository)
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
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
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
