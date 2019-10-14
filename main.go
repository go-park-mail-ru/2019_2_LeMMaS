package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/controller"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := controller.InitAPIRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
