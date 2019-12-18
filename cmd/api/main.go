package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/cmd"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/factory"
	_ "github.com/jackc/pgx/stdlib"
	"log"
	"os"
)

func main() {
	l, err := factory.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		cmd.Recover(l)
	}()

	_, err = factory.NewMiddleware()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = factory.NewMetricsHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = factory.NewAccessHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = factory.NewGameHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = factory.NewUserHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	err = factory.NewEcho().Start(":" + os.Getenv("PORT"))
	if err != nil {
		cmd.Fatal(l, err)
	}
}
