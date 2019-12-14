package delivery

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
)

type userOutput struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}

type UserUpdate struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserRegister struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
	Name     string `json:"name" valid:"required"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
}

func OutputUsers(users []*model.User) []userOutput {
	res := make([]userOutput, 0, len(users))
	for _, u := range users {
		res = append(res, OutputUser(u))
	}
	return res
}

func OutputUser(user *model.User) userOutput {
	return userOutput{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		AvatarPath: user.AvatarPath,
	}
}
