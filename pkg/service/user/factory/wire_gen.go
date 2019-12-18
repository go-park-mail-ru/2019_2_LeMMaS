// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/delivery/grpc"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/usecase"
	"github.com/jmoiron/sqlx"
	"os"
)

// Injectors from wire.go:

func NewUserHandler() (*grpc.UserHandler, error) {
	db, err := newDB()
	if err != nil {
		return nil, err
	}
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewDatabaseRepository(db, logger)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := grpc.NewUserHandler(userUsecase)
	return userHandler, nil
}

// wire.go:

func NewLogger() (logger.Logger, error) {
	sentry, err := logger.NewSentryLogger()
	if err != nil {
		return nil, err
	}
	return logger.NewCombinedLogger(sentry, logger.NewStdoutLogger()), nil
}

func newDB() (*sqlx.DB, error) {
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
