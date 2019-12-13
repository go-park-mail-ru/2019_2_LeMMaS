package http

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/api/delivery"
	"github.com/labstack/echo"
	"strconv"
	"time"
)

type UserHandler struct {
	delivery.Handler
	user   api.UserUsecase
	auth   api.AuthUsecase
	logger logger.Logger
}

func NewUserHandler(e *echo.Echo, user api.UserUsecase, auth api.AuthUsecase, logger logger.Logger) *UserHandler {
	h := UserHandler{user: user, auth: auth, logger: logger}
	e.GET(delivery.ApiV1UserListPath, h.handleUserList)
	e.GET(delivery.ApiV1UserByIDPath, h.handleUserByID)
	e.POST(delivery.ApiV1UserRegisterPath, h.handleUserRegister)
	e.POST(delivery.ApiV1UserLoginPath, h.handleUserLogin)
	e.POST(delivery.ApiV1UserLogoutPath, h.handleUserLogout)
	e.GET(delivery.ApiV1UserProfilePath, h.handleUserProfile)
	e.POST(delivery.ApiV1UserUpdatePath, h.handleUserUpdate)
	e.POST(delivery.ApiV1UserAvatarUploadPath, h.handleAvatarUpload)
	e.GET(delivery.ApiV1UserGetAvatarByNamePath, h.handleGetAvatarByName)
	return &h
}

type userToOutput struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}

func (h *UserHandler) handleUserList(c echo.Context) error {
	users, err := h.user.GetAllUsers()
	if err != nil {
		return h.Error(c, "error loading users")
	}
	usersToOutput := h.convertUsersForOutput(users)
	return h.OkWithBody(c, map[string]interface{}{
		"users": usersToOutput,
	})
}

func (h *UserHandler) handleUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.Error(c, "user id must be an integer")
	}
	userByID, err := h.user.GetUserByID(userID)
	if err != nil {
		return h.Error(c, "error loading user")
	}
	if userByID == nil {
		return h.Error(c, "user with this id not found")
	}
	return h.OkWithBody(c, map[string]interface{}{
		"user": h.convertUserForOutput(*userByID),
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

func (h *UserHandler) handleUserUpdate(c echo.Context) error {
	currentUser, err := h.currentUser(c)
	if err != nil {
		return h.Error(c, err.Error())
	}
	userToUpdate := &userToUpdate{}
	if err := c.Bind(userToUpdate); err != nil {
		return h.Error(c, "unknown error")
	}
	err = h.user.UpdateUser(currentUser.ID, userToUpdate.Password, userToUpdate.Name)
	if err != nil {
		return h.Error(c, "error updating user")
	}
	return h.Ok(c)
}

func (h *UserHandler) handleAvatarUpload(c echo.Context) error {
	currentUser, err := h.currentUser(c)
	if err != nil {
		return h.Error(c, err.Error())
	}
	err = c.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		h.logger.Error(err)
		return h.Error(c, "bad request")
	}
	avatarFile, _, err := c.Request().FormFile("avatar")
	if err != nil {
		h.logger.Error(err)
		return h.Error(c, "bad request")
	}
	defer avatarFile.Close()
	err = h.user.UpdateUserAvatar(currentUser, avatarFile)
	if err != nil {
		return h.Error(c, "error updating avatar")
	}
	return h.Ok(c)
}

func (h *UserHandler) handleGetAvatarByName(c echo.Context) error {
	name := c.FormValue("name")
	avatarUrl := h.user.GetAvatarUrlByName(name)
	return h.OkWithBody(c, map[string]string{
		"avatar_url": avatarUrl,
	})
}

func (h *UserHandler) handleUserProfile(c echo.Context) error {
	currentUser, err := h.currentUser(c)
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

func (h *UserHandler) handleUserRegister(c echo.Context) error {
	userToRegister := &userToRegister{}
	if err := c.Bind(userToRegister); err != nil {
		return h.Error(c, err.Error())
	}
	if ok, errs := h.Validate(userToRegister); !ok {
		return h.Errors(c, errs)
	}
	err := h.auth.Register(userToRegister.Email, userToRegister.Password, userToRegister.Name)
	if err != nil {
		return h.Error(c, err.Error())
	}
	return h.Ok(c)
}

type userToLogin struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
}

func (h *UserHandler) handleUserLogin(c echo.Context) error {
	userToLogin := &userToLogin{}
	if err := c.Bind(userToLogin); err != nil {
		return err
	}
	if ok, errs := h.Validate(userToLogin); !ok {
		return h.Errors(c, errs)
	}
	sessionID, err := h.auth.Login(userToLogin.Email, userToLogin.Password)
	if err != nil {
		return h.Error(c, err.Error())
	}
	h.SetCookie(c, delivery.SessionIDCookieName, sessionID, time.Now().Add(delivery.SessionIDCookieExpire))
	return h.Ok(c)
}

func (h *UserHandler) handleUserLogout(c echo.Context) error {
	sessionIDCookie, err := c.Cookie(delivery.SessionIDCookieName)
	if err != nil {
		return h.Error(c, "no session cookie")
	}
	h.DeleteCookie(c, delivery.SessionIDCookieName)
	err = h.auth.Logout(sessionIDCookie.Value)
	if err != nil {
		return h.Error(c, err.Error())
	}
	return h.Ok(c)
}

func (h *UserHandler) currentUser(c echo.Context) (*model.User, error) {
	sessionIDCookie, err := c.Cookie(delivery.SessionIDCookieName)
	if err != nil {
		return nil, errors.New("no session cookie")
	}
	currentUser, _ := h.auth.GetUserBySessionID(sessionIDCookie.Value)
	if currentUser == nil {
		return nil, errors.New("invalid session id")
	}
	return currentUser, nil
}
