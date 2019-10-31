package usecase

import (
	"crypto/md5"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user"
	"github.com/google/uuid"
	"io"
	"strings"
)

type userUsecase struct {
	userRepository     user.UserRepository
	userFileRepository user.UserFileRepository
	sessions           map[string]int
}

func NewUserUsecase(userRepository user.UserRepository, userFileRepository user.UserFileRepository) *userUsecase {
	return &userUsecase{
		userRepository:     userRepository,
		userFileRepository: userFileRepository,
		sessions:           map[string]int{},
	}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	return u.userRepository.GetAll()
}

func (u *userUsecase) GetUserBySessionID(sessionID string) (*model.User, error) {
	userID := u.sessions[sessionID]
	return u.userRepository.GetByID(userID)
}

func (u *userUsecase) UpdateUser(id int, password, name string) error {
	userToUpdate, err := u.userRepository.GetByID(id)
	if err != nil {
		return err
	}
	if password != "" {
		userToUpdate.PasswordHash = u.getPasswordHash(password)
	}
	if name != "" && userToUpdate.Name != name {
		userToUpdate.Name = name
		avatarPath := u.GetAvatarUrlByName(name)
		if avatarPath != "" {
			userToUpdate.AvatarPath = avatarPath
		}
	}
	u.userRepository.Update(userToUpdate)
	return nil
}

func (u *userUsecase) UpdateUserAvatar(user *model.User, avatarFile io.Reader, avatarPath string) error {
	newAvatarPath, err := u.userFileRepository.StoreAvatar(user, avatarFile, avatarPath)
	if err != nil {
		return err
	}
	return u.userRepository.UpdateAvatarPath(user.ID, newAvatarPath)
}

func (u *userUsecase) GetAvatarUrlByName(name string) string {
	avatarsByName := map[string]string{
		"eath":   "http://www.i2clipart.com/cliparts/3/d/1/e/clipart-earth-3d1e.png",
		"trump":  "https://www.jing.fm/clipimg/full/21-213906_trump-clipart-overload-trump-thinking-transparent.png",
		"heroku": "https://railsware.com/blog/wp-content/uploads/2017/12/How-to-set-up-the-Heroku.png",
		"cat":    "https://i.pinimg.com/originals/90/a8/56/90a856d434dd9df24d8d5fdf4bf3ce72.png",
	}
	return avatarsByName[strings.ToLower(name)]
}

func (u *userUsecase) Register(email, password, name string) error {
	userWithSameEmail, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return err
	}
	if userWithSameEmail != nil {
		return fmt.Errorf("user with email %v already registered", email)
	}
	passwordHash := u.getPasswordHash(password)
	return u.userRepository.Create(email, passwordHash, name)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	userToLogin, err := u.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if userToLogin == nil {
		return "", fmt.Errorf("incorrect email")
	}
	if u.getPasswordHash(password) != userToLogin.PasswordHash {
		return "", fmt.Errorf("incorrect password")
	}
	sessionID := u.getNewSessionID()
	u.sessions[sessionID] = userToLogin.ID
	return sessionID, nil
}

func (u *userUsecase) Logout(sessionID string) error {
	if _, ok := u.sessions[sessionID]; !ok {
		return fmt.Errorf("session id not found")
	}
	delete(u.sessions, sessionID)
	return nil
}

func (u *userUsecase) getPasswordHash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func (u *userUsecase) getNewSessionID() string {
	return uuid.New().String()
}
