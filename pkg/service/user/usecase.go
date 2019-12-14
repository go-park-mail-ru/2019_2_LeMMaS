package user

type UserUsecase interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(userID int) (*model.User, error)
	UpdateUser(userID int, password, name string) error
	UpdateUserAvatar(userID int, avatarPath string) error
	GetSpecialAvatar(name string) string
}
