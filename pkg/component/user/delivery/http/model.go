//go:generate easyjson -all

package http

type userToOutput struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	AvatarPath string `json:"avatar_path"`
}

type userToUpdate struct {
	Password string `json:"password"`
	Name     string `json:"name"`
}

type userToRegister struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
	Name     string `json:"name" valid:"required"`
}

type userToLogin struct {
	Email    string `json:"email" valid:"email,required"`
	Password string `json:"password" valid:"required"`
}
