package main

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	userHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	userRepository "github.com/go-park-mail-ru/2019_2_LeMMaS/user/repository"
	userUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/user/usecase"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
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
	e.Static("static", "static")
	http.InitMiddlewares(e)
	//db, err := getDB()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	initUserHandler(e, nil)
	log.Fatal(e.Start(":" + port))
}

func getDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initUserHandler(e *echo.Echo, db *sqlx.DB) {
	userRepo := userRepository.NewDatabaseUserRepository(db)
	userFileRepo := userRepository.NewUserFileRepository()
	usecase := userUsecase.NewUserUsecase(userRepo, userFileRepo)
	userHttpDelivery.NewUserHandler(e, usecase)
}
