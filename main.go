package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/middleware"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user/repository"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user/usecase"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	mw := middleware.NewMiddleware()
	e.Use(mw.CORS)
	e.Use(mw.Panic)

	userRepo := repository.NewMemoryUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	httpDelivery.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start("localhost:" + port))
}
