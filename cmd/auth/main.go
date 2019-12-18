package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/cmd"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/factory"
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

	h, err := factory.NewAuthHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	err = h.Serve(":" + os.Getenv("PORT"))
	if err != nil {
		cmd.Fatal(l, err)
	}
}
