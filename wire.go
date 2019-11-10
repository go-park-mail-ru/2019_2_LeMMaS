package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/business/access/delivery/http"
	"github.com/google/wire"
)

func InitAccessHandler() http.AccessHandler {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
