package cmd

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"os"
)

func Fatal(logger logger.Logger, err error) {
	logger.Error(err)
	os.Exit(1)
}

func Recover(logger logger.Logger) {
	if err := recover(); err != nil {
		logger.Errorf("%v", err)
		panic(err)
	}
}
