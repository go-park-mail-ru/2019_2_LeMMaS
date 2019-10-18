package http

import (
	baseHttp "github.com/go-park-mail-ru/2019_2_LeMMaS/internal/delivery/http"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user"
	"github.com/labstack/echo"
	"net/http"
)

const (
	ApiV1PathPrefix   = "/api/v1"
	ApiV1UserListPath = ApiV1PathPrefix + "/user/list"
	//ApiV1UserUpdatePath       = ApiV1PathPrefix + "/user/update"
	//ApiV1UserAvatarUploadPath = ApiV1PathPrefix + "/user/avatar/upload"
	//ApiV1UserProfilePath      = ApiV1PathPrefix + "/user/me"
	ApiV1UserRegisterPath = ApiV1PathPrefix + "/user/register"
	//ApiV1UserLoginPath        = ApiV1PathPrefix + "/user/login"
	//ApiV1UserLogoutPath       = ApiV1PathPrefix + "/user/logout"
)

//const (
//	SessionIDCookieName   = "session_id"
//	SessionIDCookieExpire = 10 * time.Hour
//)

type UserHandler struct {
	userUsecase user.Usecase
}

func NewUserHandler(e *echo.Echo, userUsecase user.Usecase) {
	handler := UserHandler{userUsecase: userUsecase}
	e.GET(ApiV1UserListPath, handler.HandleUserList)
	e.POST(ApiV1UserRegisterPath, handler.HandleUserRegister)
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
	return c.JSON(http.StatusOK, baseHttp.OkWithBody(map[string]interface{}{
		"users": usersToOutput,
	}))
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

//func (h *UserHandler) HandleUserUpdate(c echo.Context) error {
//currentUser, err := h.getCurrentUser(c.Request())
//if err != nil {
//	return err
//}
//decoder := json.NewDecoder(r.Body)
//user := new(userToUpdate)
//err = decoder.Decode(user)
//if err != nil {
//	h.writeError(w, err)
//	return
//}
//h.userComponent.UpdateUser(currentUser.ID, user.Password, user.Name)
//h.writeOk(w)
//}

//func (h *UserHandler) HandleAvatarUpload(w http.ResponseWriter, r *http.Request) {
//h.writeCommonHeaders(w)
//currentUser, err := h.getCurrentUser(r)
//if err != nil {
//	h.writeError(w, err)
//	return
//}
//err = r.ParseMultipartForm(32 << 20)
//if err != nil {
//	h.writeError(w, err)
//	return
//}
//avatarFile, avatarFileHeader, err := r.FormFile("avatar")
//if err != nil {
//	h.writeError(w, err)
//	return
//}
//defer avatarFile.Close()
//err = h.userComponent.UpdateUserAvatar(currentUser, avatarFile, avatarFileHeader.Filename)
//if err != nil {
//	h.writeError(w, err)
//	return
//}
//h.writeOk(w)
//}

//func (h *UserHandler) HandleUserProfile(w http.ResponseWriter, r *http.Request) {
//h.writeCommonHeaders(w)
//currentUser, err := h.getCurrentUser(r)
//if err != nil {
//	h.writeOkWithBody(w, map[string]interface{}{
//		"user": nil,
//	})
//	return
//}
//h.writeOkWithBody(w, map[string]interface{}{
//	"user": h.convertUserForOutput(*currentUser),
//})
//}

type userToRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (h *UserHandler) HandleUserRegister(c echo.Context) error {
	userToRegister := &userToRegister{}
	if err := c.Bind(userToRegister); err != nil {
		return err
	}
	err := h.userUsecase.Register(userToRegister.Email, userToRegister.Password, userToRegister.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, baseHttp.Error(err))
	}
	return c.JSON(http.StatusOK, baseHttp.Ok())
}

//type UserToLogin struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}
//
//func (h *UserHandler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	h.writeCommonHeaders(w)
//	decoder := json.NewDecoder(r.Body)
//	user := new(UserToLogin)
//	err := decoder.Decode(user)
//	if err != nil {
//		h.writeError(w, err)
//		return
//	}
//	var sessionID string
//	sessionID, err = h.userComponent.Login(user.Email, user.Password)
//	if err != nil {
//		h.writeError(w, err)
//		return
//	}
//	h.setCookie(w, SessionIDCookieName, sessionID, time.Now().Add(SessionIDCookieExpire))
//	h.writeOk(w)
//}
//
//func (h *UserHandler) HandleUserLogout(w http.ResponseWriter, r *http.Request) {
//	h.writeCommonHeaders(w)
//	sessionIDCookie, err := r.Cookie(SessionIDCookieName)
//	if err != nil {
//		h.writeError(w, fmt.Errorf("no session cookie"))
//		return
//	}
//	h.deleteCookie(w, sessionIDCookie)
//	err = h.userComponent.Logout(sessionIDCookie.Value)
//	if err != nil {
//		h.writeError(w, err)
//		return
//	}
//	h.writeOk(w)
//}
//
//func (h *UserHandler) getCurrentUser(r *http.Request) (*storage.User, error) {
//	sessionIDCookie, err := r.Cookie(SessionIDCookieName)
//	if err != nil {
//		return nil, err
//	}
//	currentUser := h.userComponent.GetUserBySessionID(sessionIDCookie.Value)
//	if currentUser == nil {
//		return nil, fmt.Errorf("invalid session id")
//	}
//	return currentUser, nil
//}
