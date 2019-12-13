package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/cmd"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/factory"
	"log"
)

func main() {
	l, err := factory.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	h, err := factory.NewAuthHandler()
	if err != nil {
		cmd.Fatal(l, err)
	}

	err = h.Serve()
	if err != nil {
		cmd.Fatal(l, err)
	}
}
