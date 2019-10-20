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
	http.InitMiddlewares(e)
	db, err := getDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	initUserHandler(e, db)
	log.Fatal(e.Start(":" + port))
}

func getDB() (*sqlx.DB, error) {
	dsn := "host=localhost port=5432 dbname=lemmas user=root password=temppassword"
	db, err := sqlx.Connect("pgx", dsn)
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
	repo := userRepository.NewDatabaseUserRepository(db)
	usecase := userUsecase.NewUserUsecase(repo)
	userHttpDelivery.NewUserHandler(e, usecase)
}
