package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	ApiV1PathPrefix           = "/api/v1"
	ApiV1UserListPath         = ApiV1PathPrefix + "/user/list"
	ApiV1UserUpdatePath       = ApiV1PathPrefix + "/user/update"
	ApiV1UserAvatarUploadPath = ApiV1PathPrefix + "/user/avatar/upload"
	ApiV1UserProfilePath      = ApiV1PathPrefix + "/user/me"
	ApiV1UserRegisterPath     = ApiV1PathPrefix + "/user/register"
	ApiV1UserLoginPath        = ApiV1PathPrefix + "/user/login"
	ApiV1UserLogoutPath       = ApiV1PathPrefix + "/user/logout"
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

	router.Use(userController.PanicMiddleware)

	return useCORS(router)
}
