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
	handler := controller.InitAPIHandler()
	http.ListenAndServe(":"+port, handler)
}
