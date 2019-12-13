package cmd

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"os"
)

func Fatal(logger logger.Logger, err error) {
	logger.Error(err)
	os.Exit(1)
}
