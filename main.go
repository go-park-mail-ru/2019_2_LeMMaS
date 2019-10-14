package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/controller"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := controller.InitAPIRouter()
	http.ListenAndServe(":"+port, r)
}
