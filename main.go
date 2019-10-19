package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	userHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	userRepository "github.com/go-park-mail-ru/2019_2_LeMMaS/user/repository"
	userUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/user/usecase"
	"github.com/labstack/echo"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	http.InitMiddlewares(e)
	initUserHandler(e)
	e.Logger.Fatal(e.Start(":" + port))
}

func initUserHandler(e *echo.Echo) {
	repo := userRepository.NewMemoryUserRepository()
	usecase := userUsecase.NewUserUsecase(repo)
	userHttpDelivery.NewUserHandler(e, usecase)
}
