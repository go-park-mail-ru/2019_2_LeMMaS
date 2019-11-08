package main

import (
	accessHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/access/delivery/http"
	accessUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/access/usecase"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	userHttpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/user/delivery/http"
	userRepository "github.com/go-park-mail-ru/2019_2_LeMMaS/user/repository"
	userUsecase "github.com/go-park-mail-ru/2019_2_LeMMaS/user/usecase"
	"github.com/gomodule/redigo/redis"
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
		return
	}
	redisConn, err := getRedis()
	if err != nil {
		return
	}

	initAccessHandler(e)
	initUserHandler(e, db, redisConn)

	err = e.Start(":" + port)
	if err != nil {
		logger.Error(err)
	}
}

func getDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		logger.Errorf("cannot connect to postgres: %s", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Errorf("cannot connect to postgres: %s", err)
		return nil, err
	}
	return db, nil
}

func initAccessHandler(e *echo.Echo) {
	csrfUsecase := accessUsecase.NewCSRFUsecase()
	accessHttpDelivery.NewAccessHandler(e, csrfUsecase)
}

func getRedis() (redis.Conn, error) {
	connection, err := redis.DialURL(os.Getenv("REDIS_DSN"))
	if err != nil {
		logger.Errorf("cannot connect to redis: %s", err)
		return nil, err
	}
	_, err = connection.Do("PING")
	if err != nil {
		logger.Errorf("cannot connect to redis: %s", err)
		return nil, err
	}
	return connection, nil
}

func initUserHandler(e *echo.Echo, db *sqlx.DB, redisConn redis.Conn) {
	dbRepo := userRepository.NewDatabaseRepository(db)
	fileRepo := userRepository.NewFileRepository()
	sessionRepo := userRepository.NewSessionRepository(redisConn)
	usecase := userUsecase.NewUserUsecase(dbRepo, fileRepo, sessionRepo)
	userHttpDelivery.NewUserHandler(e, usecase)
}
