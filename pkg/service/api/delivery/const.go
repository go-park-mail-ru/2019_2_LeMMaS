package delivery

import (
	"strings"
	"time"
)

const (
	publicPrefix  = "/api/v1/public"
	privatePrefix = "/api/v1/private"

	ApiV1AccessCSRFPath = publicPrefix + "/access/csrf"

	ApiV1GamePath = privatePrefix + "/game"

	ApiV1UserListPath            = publicPrefix + "/user/list"
	ApiV1UserByIDPath            = publicPrefix + "/user/:id"
	ApiV1UserRegisterPath        = publicPrefix + "/user/register"
	ApiV1UserLoginPath           = publicPrefix + "/user/login"
	ApiV1UserLogoutPath          = privatePrefix + "/user/logout"
	ApiV1UserProfilePath         = privatePrefix + "/user/me"
	ApiV1UserUpdatePath          = privatePrefix + "/user/update"
	ApiV1UserAvatarUploadPath    = privatePrefix + "/user/avatar/upload"
	ApiV1UserGetAvatarByNamePath = privatePrefix + "/user/avatar/getByName"
)

const (
	MetricsPath = "/metrics"
)

const (
	SessionCookieName   = "session_id"
	SessionCookieExpire = 10 * time.Hour
)

func IsPrivatePath(path string) bool {
	return strings.HasPrefix(path, privatePrefix)
}
