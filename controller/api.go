package controller

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	ApiV1UserListPath         = "/api/v1/user/list"
	ApiV1UserUpdatePath       = "/api/v1/user/update"
	ApiV1UserAvatarUploadPath = "/api/v1/user/avatar/upload"
	ApiV1UserProfilePath      = "/api/v1/user/me"
	ApiV1UserRegisterPath     = "/api/v1/user/register"
	ApiV1UserLoginPath        = "/api/v1/user/login"
	ApiV1UserLogoutPath       = "/api/v1/user/logout"
)

func InitAPIHandler() http.Handler {
	router := mux.NewRouter()

	userController := NewUserController()
	router.HandleFunc(ApiV1UserListPath, userController.HandleUserList).Methods(http.MethodGet)
	router.HandleFunc(ApiV1UserUpdatePath, userController.HandleUserUpdate).Methods(http.MethodPost)
	router.HandleFunc(ApiV1UserAvatarUploadPath, userController.HandleAvatarUpload).Methods(http.MethodPost)
	router.HandleFunc(ApiV1UserProfilePath, userController.HandleUserProfile).Methods(http.MethodGet)
	router.HandleFunc(ApiV1UserRegisterPath, userController.HandleUserRegister).Methods(http.MethodPost)
	router.HandleFunc(ApiV1UserLoginPath, userController.HandleUserLogin).Methods(http.MethodPost)
	router.HandleFunc(ApiV1UserLogoutPath, userController.HandleUserLogout).Methods(http.MethodPost)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"localhost", ".now.sh"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}
