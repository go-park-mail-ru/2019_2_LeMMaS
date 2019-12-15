//+build wireinject
//go:generate wire

package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	delivery "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game/delivery/grpc"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/game/usecase"
	"github.com/google/wire"
)

func NewGameHandler() (*delivery.GameHandler, error) {
	wire.Build(
		delivery.NewGameHandler,
		usecase.NewGameUsecase,
		repository.NewRepository,
		NewLogger,
	)
	return &delivery.GameHandler{}, nil
}

func NewLogger() (logger.Logger, error) {
	sentry, err := logger.NewSentryLogger()
	if err != nil {
		return nil, err
	}
	return logger.NewCombinedLogger(sentry, logger.NewStdoutLogger()), nil
}
