package http

import (
	"strings"
	"time"
)

const (
	// API
	apiPublic  = "/api/v1/public"
	apiPrivate = "/api/v1/private"

	ApiV1AccessCSRFPath = apiPublic + "/access/csrf"

	ApiV1GamePath = apiPrivate + "/game"

	ApiV1UserListPath            = apiPublic + "/user/list"
	ApiV1UserByIDPath            = apiPublic + "/user/:id"
	ApiV1UserRegisterPath        = apiPublic + "/user/register"
	ApiV1UserLoginPath           = apiPublic + "/user/login"
	ApiV1UserLogoutPath          = apiPrivate + "/user/logout"
	ApiV1UserProfilePath         = apiPrivate + "/user/me"
	ApiV1UserUpdatePath          = apiPrivate + "/user/update"
	ApiV1UserAvatarUploadPath    = apiPrivate + "/user/avatar/upload"
	ApiV1UserGetAvatarByNamePath = apiPrivate + "/user/avatar/getByName"

	// Support
	MetricsPath = "/metrics"
)

const (
	SessionIDCookieName   = "session_id"
	SessionIDCookieExpire = 10 * time.Hour
)

func IsPrivatePath(path string) bool {
	return strings.HasPrefix(path, apiPrivate)
}
