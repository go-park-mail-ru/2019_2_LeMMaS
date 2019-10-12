package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/handlers"
	"log"
	"net/http"
)

var (
	apiPath = "/api/v1"
	PORT    = ":8080"
)

func main() {
	//mux := http.NewServeMux() // TODO сделать через мультиплексор

	//mux.HandleFunc()
	// TODO парсить из командной строки параметры конфигурации
	loginHandler := http.HandlerFunc(handlers.LoginHandler)
	logoutHandler := http.HandlerFunc(handlers.LogoutHandler)
	registerHandler := http.HandlerFunc(handlers.RegisterHandler)
	userDataHandler := http.HandlerFunc(handlers.GetUserDataHandler)
	changeUserDataHandler := http.HandlerFunc(handlers.ChangeUserDataHandler)
	uploadAvatarHandler := http.HandlerFunc(handlers.UploadAvatarHandler)

	http.Handle(apiPath+"/login", handlers.MethodMiddleware("POST")(loginHandler))
	http.Handle(apiPath+"/logout", handlers.MethodMiddleware("POST")(logoutHandler))
	http.Handle(apiPath+"/register", handlers.MethodMiddleware("POST")(registerHandler))
	http.Handle(apiPath+"/user", handlers.MethodMiddleware("GET")(userDataHandler))
	http.Handle(apiPath+"/user/upload", handlers.MethodMiddleware("PUT")(uploadAvatarHandler))
	http.Handle(apiPath+"/user/settings", handlers.MethodMiddleware("PATCH")(changeUserDataHandler))

	fmt.Printf("starting server at %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
