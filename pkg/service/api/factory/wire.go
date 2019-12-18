//+build wireinject
//go:generate wire

package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery/ws"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/google/wire"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"os"
)

func NewMiddleware() (http.Middleware, error) {
	wire.Build(
		http.NewMiddleware,
		api.NewMetrics,
		NewEcho,
		NewLogger,
	)
	return http.Middleware{}, nil
}

func NewMetricsHandler() (*http.MetricsHandler, error) {
	wire.Build(
		http.NewMetricsHandler,
		NewEcho,
	)
	return &http.MetricsHandler{}, nil
}

func NewAccessHandler() (*http.AccessHandler, error) {
	wire.Build(
		http.NewAccessHandler,
		usecase.NewCSRFUsecase,
		NewEcho,
		NewLogger,
	)
	return &http.AccessHandler{}, nil
}

func NewGameHandler() (*ws.GameHandler, error) {
	wire.Build(
		ws.NewGameHandler,
		usecase.NewGameUsecase,
		usecase.NewAuthUsecase,
		newGameClient,
		newAuthClient,
		NewEcho,
		NewLogger,
	)
	return &ws.GameHandler{}, nil
}

func NewUserHandler() (*http.UserHandler, error) {
	wire.Build(
		http.NewUserHandler,
		usecase.NewUserUsecase,
		usecase.NewAuthUsecase,
		repository.NewS3Repository,
		newUserClient,
		newAuthClient,
		NewEcho,
		NewLogger,
	)
	return &http.UserHandler{}, nil
}

var echoInstance *echo.Echo = nil

func NewEcho() *echo.Echo {
	if echoInstance == nil {
		echoInstance = echo.New()
	}
	return echoInstance
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

func newAuthClient() (auth.AuthClient, error) {
	conn, err := newGRPC("auth:" + os.Getenv("PORT"))
	return auth.NewAuthClient(conn), err
}

func newUserClient() (user.UserClient, error) {
	conn, err := newGRPC("user:" + os.Getenv("PORT"))
	return user.NewUserClient(conn), err
}

func newGameClient() (game.GameClient, error) {
	conn, err := newGRPC("game:" + os.Getenv("PORT"))
	return game.NewGameClient(conn), err
}

func newGRPC(url string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
