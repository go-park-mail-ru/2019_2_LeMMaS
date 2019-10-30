package main

import (
	//"2019_2_LeMMaS/logger"
	"github.com/getsentry/sentry-go"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	// logger "github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
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
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://23f2e9b5b60c448c87463bc49ffc8396@sentry.io/1801031",
	})
	// logger.InitLogger()
	db, err := getDB()

	if err != nil {
		log.Fatal(err)
		return
	}
	initUserHandler(e, db)
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
