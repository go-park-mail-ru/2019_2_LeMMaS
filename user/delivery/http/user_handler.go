package http

import (
	"fmt"
	httpDelivery "github.com/go-park-mail-ru/2019_2_LeMMaS/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user"
	"github.com/labstack/echo"
	"time"
)

const (
	ApiV1PathPrefix = "/api/v1"

	ApiV1UserListPath         = ApiV1PathPrefix + "/user/list"
	ApiV1UserUpdatePath       = ApiV1PathPrefix + "/user/update"
	ApiV1UserAvatarUploadPath = ApiV1PathPrefix + "/user/avatar/upload"
	ApiV1UserProfilePath      = ApiV1PathPrefix + "/user/me"
	ApiV1UserRegisterPath     = ApiV1PathPrefix + "/user/register"
	ApiV1UserLoginPath        = ApiV1PathPrefix + "/user/login"
	ApiV1UserLogoutPath       = ApiV1PathPrefix + "/user/logout"
)

const (
	SessionIDCookieName   = "session_id"
	SessionIDCookieExpire = 10 * time.Hour
)

type UserHandler struct {
	userUsecase user.Usecase
	httpDelivery.Handler
}

func NewUserHandler(e *echo.Echo, userUsecase user.Usecase) *UserHandler {
	handler := UserHandler{userUsecase: userUsecase}
	e.GET(ApiV1UserListPath, handler.HandleUserList)
	e.POST(ApiV1UserRegisterPath, handler.HandleUserRegister)
	e.POST(ApiV1UserLoginPath, handler.HandleUserLogin)
	e.POST(ApiV1UserLogoutPath, handler.HandleUserLogout)
	e.GET(ApiV1UserProfilePath, handler.HandleUserProfile)
	e.POST(ApiV1UserUpdatePath, handler.HandleUserUpdate)
	e.POST(ApiV1UserAvatarUploadPath, handler.HandleAvatarUpload)
	return &handler
}

type userToOutput struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}

func (h *UserHandler) HandleUserList(c echo.Context) error {
	users := h.userUsecase.GetAllUsers()
	usersToOutput := h.convertUsersForOutput(users)
	return h.OkWithBody(c, map[string]interface{}{
		"users": usersToOutput,
	})
}

func (h *UserHandler) convertUsersForOutput(users []model.User) []userToOutput {
	usersToOutput := make([]userToOutput, 0, len(users))
	for _, u := range users {
		usersToOutput = append(usersToOutput, h.convertUserForOutput(u))
	}
	return usersToOutput
}

func (h *UserHandler) convertUserForOutput(user model.User) userToOutput {
	return userToOutput{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		AvatarPath: user.AvatarPath,
	}
}

type userToUpdate struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (h *UserHandler) HandleUserUpdate(c echo.Context) error {
	currentUser, err := h.getCurrentUser(c)
	if err != nil {
		return h.Error(c, err)
	}
	userToUpdate := &userToUpdate{}
	if err := c.Bind(userToUpdate); err != nil {
		return h.Error(c, err)
	}
	h.userUsecase.UpdateUser(currentUser.ID, userToUpdate.Password, userToUpdate.Name)
	return h.Ok(c)
}

func (h *UserHandler) HandleAvatarUpload(c echo.Context) error {
	currentUser, err := h.getCurrentUser(c)
	if err != nil {
		return h.Error(c, err)
	}
	err = c.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		return h.Error(c, err)
	}
	avatarFile, avatarFileHeader, err := c.Request().FormFile("avatar")
	if err != nil {
		return h.Error(c, err)
	}
	defer avatarFile.Close()
	err = h.userUsecase.UpdateUserAvatar(currentUser, avatarFile, avatarFileHeader.Filename)
	if err != nil {
		return h.Error(c, err)
	}
	return h.Ok(c)
}

func (h *UserHandler) HandleUserProfile(c echo.Context) error {
	currentUser, err := h.getCurrentUser(c)
	if err != nil {
		return h.OkWithBody(c, map[string]interface{}{
			"user": nil,
		})
	}
	return h.OkWithBody(c, map[string]interface{}{
		"user": h.convertUserForOutput(*currentUser),
	})
}

type userToRegister struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
	Name     string `json:"name" valid:"required"`
}

func (h *UserHandler) HandleUserRegister(c echo.Context) error {
	userToRegister := &userToRegister{}
	if err := c.Bind(userToRegister); err != nil {
		return err
	}
	if ok, errors := h.Validate(userToRegister); !ok {
		return h.Errors(c, errors)
	}
	err := h.userUsecase.Register(userToRegister.Email, userToRegister.Password, userToRegister.Name)
	if err != nil {
		return h.Error(c, err)
	}
	return h.Ok(c)
}

type userToLogin struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
}

func (h *UserHandler) HandleUserLogin(c echo.Context) error {
	userToLogin := &userToLogin{}
	if err := c.Bind(userToLogin); err != nil {
		return err
	}
	if ok, errors := h.Validate(userToLogin); !ok {
		return h.Errors(c, errors)
	}
	sessionID, err := h.userUsecase.Login(userToLogin.Email, userToLogin.Password)
	if err != nil {
		return h.Error(c, err)
	}
	h.SetCookie(c, SessionIDCookieName, sessionID, time.Now().Add(SessionIDCookieExpire))
	return h.Ok(c)
}

func (h *UserHandler) HandleUserLogout(c echo.Context) error {
	sessionIDCookie, err := c.Cookie(SessionIDCookieName)
	if err != nil {
		return h.Error(c, fmt.Errorf("no session cookie"))
	}
	h.DeleteCookie(c, SessionIDCookieName)
	err = h.userUsecase.Logout(sessionIDCookie.Value)
	if err != nil {
		return h.Error(c, err)
	}
	return h.Ok(c)
}

func (h *UserHandler) getCurrentUser(c echo.Context) (*model.User, error) {
	sessionIDCookie, err := c.Cookie(SessionIDCookieName)
	if err != nil {
		return nil, err
	}
	currentUser := h.userUsecase.GetUserBySessionID(sessionIDCookie.Value)
	if currentUser == nil {
		return nil, fmt.Errorf("invalid session id")
	}
	return currentUser, nil
}
