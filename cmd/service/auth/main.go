package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/cmd"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"log"
)

func main() {
	l, err := api.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	err = api.NewEcho().Start(":" + cmd.GetPort())
	if err != nil {
		cmd.Fatal(l, err)
	}
}
