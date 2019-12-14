package factory

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/delivery/grpc"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"os"
)

func NewUserHandler() (grpc.UserHandler, error) {
	wire.Build(
		grpc.NewUserHandler,
		NewLogger,
	)
	return grpc.UserHandler{}, nil
}

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
