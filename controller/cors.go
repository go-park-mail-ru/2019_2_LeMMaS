package controller

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
)

func useCORS(router *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOriginValidator(func(origin string) bool {
		isNowSh, _ := regexp.MatchString(`^https:\/\/20192lemmas-.*\.now\.sh$`, origin)
		isLocalhost, _ := regexp.MatchString(`^http:\/\/localhost:\d*$`, origin)
		return isNowSh || isLocalhost
	})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	credentials := handlers.AllowCredentials()
	return handlers.CORS(originsOk, headersOk, methodsOk, credentials)(router)
}
