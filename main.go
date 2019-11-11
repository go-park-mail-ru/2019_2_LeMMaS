package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/logger"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"os"
)

func main() {
	l, err := NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	_, err = InitMiddleware()
	if err != nil {
		fatal(l, err)
	}
	_, err = InitAccessHandler()
	if err != nil {
		fatal(l, err)
	}
	_, err = InitUserHandler()
	if err != nil {
		fatal(l, err)
	}
	err = NewEcho().Start(":" + getPort())
	if err != nil {
		fatal(l, err)
	}
}

func fatal(logger logger.Logger, err error) {
	logger.Error(err)
	os.Exit(1)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
