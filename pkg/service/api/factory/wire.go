//+build wireinject
//go:generate wire

package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery/ws"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/google/wire"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"os"
)

func NewMiddleware() (http.Middleware, error) {
	wire.Build(
		http.NewMiddleware,
		NewEcho,
		NewLogger,
	)
	return http.Middleware{}, nil
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
		NewAuthClient,
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
		NewUserClient,
		NewAuthClient,
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

func NewAuthClient() (auth.AuthClient, error) {
	conn, err := newGRPC(os.Getenv("AUTH_URL"))
	return auth.NewAuthClient(conn), err
}

func NewUserClient() (user.UserClient, error) {
	conn, err := newGRPC(os.Getenv("USER_URL"))
	return user.NewUserClient(conn), err
}

func newGRPC(url string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
