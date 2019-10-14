package controller

import (
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

func InitAPIRouter() *mux.Router {
	r := mux.NewRouter()

	userController := NewUserController()
	r.HandleFunc(ApiV1UserListPath, userController.HandleUserList).Methods(http.MethodGet)
	r.HandleFunc(ApiV1UserUpdatePath, userController.HandleUserUpdate).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserAvatarUploadPath, userController.HandleAvatarUpload).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserProfilePath, userController.HandleUserProfile).Methods(http.MethodGet)
	r.HandleFunc(ApiV1UserRegisterPath, userController.HandleUserRegister).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserLoginPath, userController.HandleUserLogin).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserLogoutPath, userController.HandleUserLogout).Methods(http.MethodPost)

	return r
}
