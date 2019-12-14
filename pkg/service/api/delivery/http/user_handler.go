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

func (h *UserHandler) handleUserList(c echo.Context) error {
	users, err := h.user.GetAll()
	if err != nil {
		return h.Error(c, "error loading users")
	}
	return h.OkWithBody(c, map[string]interface{}{
		"users": delivery.OutputUsers(users),
	})
}

func (h *UserHandler) handleUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.Error(c, "user id must be an integer")
	}
	userByID, err := h.user.GetByID(userID)
	if err != nil {
		return h.Error(c, "error loading user")
	}
	if userByID == nil {
		return h.Error(c, "user with this id not found")
	}
	return h.OkWithBody(c, map[string]interface{}{
		"user": delivery.OutputUser(userByID),
	})
}

func (h *UserHandler) handleUserUpdate(c echo.Context) error {
	currentUser, err := h.currentUser(c)
	if err != nil {
		return h.Error(c, err.Error())
	}
	u := &delivery.UserUpdate{}
	if err := c.Bind(u); err != nil {
		return h.Error(c, "unknown error")
	}
	err = h.user.Update(currentUser.ID, u.Password, u.Name)
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
	err = h.user.UpdateAvatar(currentUser.ID, avatarFile)
	if err != nil {
		return h.Error(c, "error updating avatar")
	}
	return h.Ok(c)
}

func (h *UserHandler) handleGetAvatarByName(c echo.Context) error {
	name := c.FormValue("name")
	avatarUrl := h.user.GetSpecialAvatar(name)
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
		"user": delivery.OutputUser(currentUser),
	})
}

func (h *UserHandler) handleUserRegister(c echo.Context) error {
	u := &delivery.UserRegister{}
	if err := c.Bind(u); err != nil {
		return h.Error(c, err.Error())
	}
	if ok, errs := h.Validate(u); !ok {
		return h.Errors(c, errs)
	}
	err := h.auth.Register(u.Email, u.Password, u.Name)
	if err != nil {
		return h.Error(c, err.Error())
	}
	return h.Ok(c)
}

func (h *UserHandler) handleUserLogin(c echo.Context) error {
	u := &delivery.UserLogin{}
	if err := c.Bind(u); err != nil {
		return err
	}
	if ok, errs := h.Validate(u); !ok {
		return h.Errors(c, errs)
	}
	sessionID, err := h.auth.Login(u.Email, u.Password)
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
	currentUser, _ := h.auth.GetUserBySession(sessionIDCookie.Value)
	if currentUser == nil {
		return nil, errors.New("invalid session id")
	}
	return currentUser, nil
}
