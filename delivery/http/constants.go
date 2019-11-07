package http

import "time"

const (
	ApiV1Public  = "/api/v1/public"
	ApiV1Private = "/api/v1/private"

	ApiV1UserListPath            = ApiV1Public + "/user/list"
	ApiV1UserRegisterPath        = ApiV1Public + "/user/register"
	ApiV1UserLoginPath           = ApiV1Public + "/user/login"
	ApiV1UserLogoutPath          = ApiV1Private + "/user/logout"
	ApiV1UserProfilePath         = ApiV1Private + "/user/me"
	ApiV1UserUpdatePath          = ApiV1Private + "/user/update"
	ApiV1UserAvatarUploadPath    = ApiV1Private + "/user/avatar/upload"
	ApiV1UserGetAvatarByNamePath = ApiV1Private + "/user/avatar/getByName"

	ApiV1AccessCSRFPath = ApiV1Public + "/access/csrf"
)

const (
	SessionIDCookieName   = "session_id"
	SessionIDCookieExpire = 10 * time.Hour
)
