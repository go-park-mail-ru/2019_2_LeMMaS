package handlers

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/config"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/cookie"
	db "github.com/go-park-mail-ru/2019_2_LeMMaS/database"
	"io"
	"net/http"
	"os"
)

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func MethodMiddleware(method string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				w.WriteHeader(405)
				_, _ = w.Write([]byte("Method not allowed"))
			}
			next.ServeHTTP(w, r)
		})
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	decoder := json.NewDecoder(r.Body)
	var curUser config.AuthConfig
	err := decoder.Decode(&curUser)
	if err != nil {
		panic(err)
	}
	if !db.IsUserAuthCorrect(curUser.Login, curUser.Password) {
		w.WriteHeader(400)
		return
	}
	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	cookie.SetUserCookie(w, c, curUser.Login)
	w.WriteHeader(200)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	curCookie, err := r.Cookie("sessionId")
	if err == http.ErrNoCookie || curCookie == nil{
		w.WriteHeader(403) // если куки не найдены
		return
	}
	cookie.DeleteCookie(w, *curCookie)
	w.WriteHeader(200)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	decoder := json.NewDecoder(r.Body)
	var curUser config.AuthConfig
	err := decoder.Decode(&curUser)
	if err != nil {
		panic(err)
	}
	err = db.CreateNewUser(curUser) // TODO печатать есть ли юзер с таким логином или email
	if err != nil {
		w.WriteHeader(409) // есть пользователь с таким email или логином
		return
	}
	c := config.SessionConfig{
		Name:            "sessionId",
		LifetimeSeconds: 86400,
		Path:			 "/",
		Secure:          true,
		HTTPOnly:        true,
	}
	cookie.SetUserCookie(w, c, curUser.Login)
	w.WriteHeader(200)
}

func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	curUser := getUser(w, r)
	var nullUser db.User
	if curUser == nullUser {
		return
	}
	w.WriteHeader(200)
	res, err := json.Marshal(curUser)
	if err != nil {
		panic(err)
	}
	_, _  = w.Write(res)
}

func UploadAvatarHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	curUser := getUser(w, r)
	var nullUser db.User
	if curUser == nullUser {
		return
	}
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}
	file, fileHandler, err := r.FormFile("avatar")
	if err != nil {
		panic(err) // здесь может отправлять ответ с определенным заголовком?
	}
	defer file.Close()
	// TODO удалять старые аватарки пользователя
	f, err := os.OpenFile(db.PathAvatar + curUser.Login + fileHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)
	curUser.AvatarAddress = db.PathAvatar + curUser.Login // директория, где хранится файл с аватаром юзера
	w.WriteHeader(200)
}

func ChangeUserDataHandler(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	curUser := getUser(w, r)
	var nullUser db.User
	if curUser == nullUser {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var newUserData, nullUserData db.User
	err := decoder.Decode(&newUserData)
	if err != nil {
		panic(err)
	}
	if newUserData == nullUserData {
		w.WriteHeader(400)
		return
	}
	db.ChangeUserData(curUser, newUserData)
	w.WriteHeader(200)
}

func getUser(w http.ResponseWriter, r *http.Request) db.User { // TODO убрать возвраты заголовков (сейчас в вызывающей функции стоит return)
	curCookie, err := r.Cookie("sessionId")
	var nullUser db.User
	if err != nil || curCookie == nil || !cookie.IsInDB(*curCookie) {
		w.WriteHeader(403)
		return nullUser
	}
	curUser := db.GetUserByCookie(curCookie.Value)
	if curUser == nullUser {
		w.WriteHeader(404) // возвращать ошибку юзер не найден 404?
		return nullUser
	}
	return curUser
}