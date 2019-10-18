package http

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(corsMiddleware)
	e.Use(panicMiddleware)
}

func corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		//	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
		//originsOk := handlers.AllowedOriginValidator(func(origin string) bool {
		//	isNowSh, _ := regexp.MatchString(`^https:\/\/20192lemmas-.*\.now\.sh$`, origin)
		//	isLocalhost, _ := regexp.MatchString(`^http:\/\/localhost:\d*$`, origin)
		//	return isNowSh || isLocalhost
		//})
		//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
		//credentials := handlers.AllowCredentials()
		//return handlers.CORS(originsOk, headersOk, methodsOk, credentials)(router)

		return next(c)
	}
}

func panicMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic during request to %s: %s", c.Request().URL.Path, err)
				c.JSON(http.StatusInternalServerError, "internal error")
			}
		}()
		return next(c)
	}
}
