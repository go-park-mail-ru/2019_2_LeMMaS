package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/cmd"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/init"
	_ "github.com/jackc/pgx/stdlib"
	"log"
)

func main() {
	l, err := init.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	_, err = init.NewMiddleware()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = init.NewAccessHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = init.NewGameHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	_, err = init.NewUserHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	err = init.NewEcho().Start(":" + cmd.GetPort())
	if err != nil {
		cmd.Fatal(l, err)
	}
}
