package db

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/config"
)

var (
	users AllUsers // TODO возможно придется создать функцию инициализации нового списка юзеров
)

func CreateNewUser(curUser config.AuthConfig) error {
	if isInEmail(curUser.Email) || isInLogin(curUser.Login) {
		return  errors.New("user is already have") // TODO отправить нормальную ошибку
	}
	users.mu.Lock()
	user := User{
		Id:            len(users.Users),
		Login:         curUser.Login,
		Password:      curUser.Password,
		Email:		   curUser.Email,
		Role:          "activeUser",
	}
	users.Users = append(users.Users, user)
	users.mu.Unlock()
	return nil
}

func ChangeUserData(curUser User, newDataUser User) {
	for _, user := range users.Users {
		if user.Cookie == curUser.Cookie {
			if newDataUser.Login != "" && user.Login != newDataUser.Login {
				user.Login = newDataUser.Login
			}
			if newDataUser.Password != "" && user.Password != newDataUser.Password {
				user.Password = newDataUser.Password
			}
			if newDataUser.Email != "" && user.Email != newDataUser.Email {
				user.Email = newDataUser.Email
			}
		}
	}
}

func GetUserByCookie(cookieValue string) User {
	var nullUser User
	for _, user := range users.Users {
		if user.Cookie == cookieValue {
			return user
		}
	}
	return nullUser
}

func isInLogin(login string) bool {
	for _, user := range users.Users {
		if user.Login == login {
			return true
		}
	}
	return false
}

func isInEmail(email string) bool {
	for _, user := range users.Users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func IsUserAuthCorrect(login string, password string) bool {
	for _, user := range users.Users {
		if user.Login == login && user.Password == password {
			return true
		}
	}
	return false
}



