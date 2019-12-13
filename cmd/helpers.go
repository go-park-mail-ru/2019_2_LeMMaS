package cmd

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func Fatal(logger logger.Logger, err error) {
	logger.Error(err)
	os.Exit(1)
}
