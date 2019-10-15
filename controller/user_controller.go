package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/component"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/storage"
	"net/http"
	"time"
)

const (
	SessionIDCookieName   = "session_id"
	SessionIDCookieExpire = 10 * time.Hour
)

type UserController struct {
	Controller
	userComponent *component.UserComponent
}

func NewUserController() *UserController {
	return &UserController{
		userComponent: component.NewUserComponent(),
	}
}

type UserToOutput struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}

func (c UserController) HandleUserList(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	users := c.userComponent.GetAllUsers()
	usersToOutput := c.convertUsersForOutput(users)
	c.writeOkWithBody(w, map[string]interface{}{
		"users": usersToOutput,
	})
}

func (c UserController) convertUsersForOutput(users []storage.User) []UserToOutput {
	usersToOutput := make([]UserToOutput, 0, len(users))
	for _, user := range users {
		usersToOutput = append(usersToOutput, c.convertUserForOutput(user))
	}
	return usersToOutput
}

func (c UserController) convertUserForOutput(user storage.User) UserToOutput {
	return UserToOutput{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		AvatarPath: user.AvatarPath,
	}
}

type UserToUpdate struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c UserController) HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c.writeCommonHeaders(w)
	currentUser, err := c.getCurrentUser(r)
	if err != nil {
		c.writeError(w, err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	user := new(UserToUpdate)
	err = decoder.Decode(user)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.userComponent.UpdateUser(currentUser.ID, user.Password, user.Name)
	c.writeOk(w)
}

func (c UserController) HandleAvatarUpload(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	currentUser, err := c.getCurrentUser(r)
	if err != nil {
		c.writeError(w, err)
		return
	}
	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		c.writeError(w, err)
		return
	}
	avatarFile, avatarFileHeader, err := r.FormFile("avatar")
	if err != nil {
		c.writeError(w, err)
		return
	}
	defer avatarFile.Close()
	err = c.userComponent.UpdateUserAvatar(currentUser, avatarFile, avatarFileHeader.Filename)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.writeOk(w)
}

func (c UserController) HandleUserProfile(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	currentUser, err := c.getCurrentUser(r)
	if err != nil {
		c.writeOkWithBody(w, map[string]interface{}{
			"user": nil,
		})
		return
	}
	c.writeOkWithBody(w, map[string]interface{}{
		"user": c.convertUserForOutput(*currentUser),
	})
}

type UserToRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c UserController) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c.writeCommonHeaders(w)
	decoder := json.NewDecoder(r.Body)
	user := new(UserToRegister)
	err := decoder.Decode(user)
	if err != nil {
		c.writeError(w, err)
		return
	}
	err = c.userComponent.Register(user.Email, user.Password, user.Name)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.writeOk(w)
}

type UserToLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c UserController) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	c.writeCommonHeaders(w)
	decoder := json.NewDecoder(r.Body)
	user := new(UserToLogin)
	err := decoder.Decode(user)
	if err != nil {
		c.writeError(w, err)
		return
	}
	var sessionID string
	sessionID, err = c.userComponent.Login(user.Email, user.Password)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.setCookie(w, SessionIDCookieName, sessionID, time.Now().Add(SessionIDCookieExpire))
	c.writeOk(w)
}

func (c UserController) HandleUserLogout(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	sessionIDCookie, err := r.Cookie(SessionIDCookieName)
	if err != nil {
		c.writeError(w, fmt.Errorf("no session cookie"))
		return
	}
	c.deleteCookie(w, sessionIDCookie)
	err = c.userComponent.Logout(sessionIDCookie.Value)
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.writeOk(w)
}

func (c UserController) getCurrentUser(r *http.Request) (*storage.User, error) {
	sessionIDCookie, err := r.Cookie(SessionIDCookieName)
	if err != nil {
		return nil, err
	}
	currentUser := c.userComponent.GetUserBySessionID(sessionIDCookie.Value)
	if currentUser == nil {
		return nil, fmt.Errorf("invalid session id")
	}
	return currentUser, nil
}
