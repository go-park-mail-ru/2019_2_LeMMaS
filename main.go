package main

import (
	accessHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/access/delivery/http"
	accessUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/access/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	userHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	userRepository "github.com/go-park-mail-ru/2019_2_LeMMaS/user/repository"
	userUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/user/usecase"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	logger.Init(e)
	http.InitMiddlewares(e)
	db, err := getDB()
	if err != nil {
		logger.Error(err)
		return
	}

	initAccessHandler(e)
	initUserHandler(e, db)

	e.Logger.Fatal(e.Start(":" + port))
	err = e.Start(":" + port)
	if err != nil {
		logger.Error(err)
	}
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

func initAccessHandler(e *echo.Echo) {
	csrfUsecase := accessUsecase.NewCSRFUsecase()
	accessHttpDelivery.NewAccessHandler(e, csrfUsecase)
}

func initUserHandler(e *echo.Echo, db *sqlx.DB) {
	userRepo := userRepository.NewDatabaseUserRepository(db)
	userFileRepo := userRepository.NewUserFileRepository()
	usecase := userUsecase.NewUserUsecase(userRepo, userFileRepo)
	userHttpDelivery.NewUserHandler(e, usecase)
}
