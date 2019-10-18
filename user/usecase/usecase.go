package usecase

import (
	"crypto/md5"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/user"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

const UserAvatarDirectory = "_storage/user/avatar"

type userUsecase struct {
	userRepository user.Repository
	sessions       map[string]int
}

func NewUserUsecase(userRepository user.Repository) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
		sessions:       map[string]int{},
	}
}

func (u userUsecase) GetAllUsers() []model.User {
	return u.userRepository.GetAll()
}

func (u userUsecase) GetUserBySessionID(sessionID string) *model.User {
	userID := u.sessions[sessionID]
	return u.userRepository.GetByID(userID)
}

func (u userUsecase) UpdateUser(id int, password, name string) {
	passwordHash := ""
	if password != "" {
		passwordHash = u.getPasswordHash(password)
	}
	u.userRepository.Update(id, passwordHash, name)
}

func (u userUsecase) UpdateUserAvatar(user *model.User, avatarFile io.Reader, avatarPath string) error {
	u.deleteFileIfExists(user.AvatarPath)
	newAvatarPath, err := u.storeUserAvatar(user.ID, avatarFile, avatarPath)
	if err != nil {
		return err
	}
	u.userRepository.UpdateAvatarPath(user.ID, newAvatarPath)
	return nil
}

func (u userUsecase) Register(email, password, name string) error {
	if u.userRepository.GetByEmail(email) != nil {
		return fmt.Errorf("user with email %v already registered", email)
	}
	passwordHash := u.getPasswordHash(password)
	u.userRepository.Create(email, passwordHash, name)
	return nil
}

func (u userUsecase) Login(email, password string) (string, error) {
	userToLogin := u.userRepository.GetByEmail(email)
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

func (u userUsecase) Logout(sessionID string) error {
	if _, ok := u.sessions[sessionID]; !ok {
		return fmt.Errorf("session id not found")
	}
	delete(u.sessions, sessionID)
	return nil
}

func (u userUsecase) getPasswordHash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func (u userUsecase) getNewSessionID() string {
	return uuid.New().String()
}

func (u userUsecase) storeUserAvatar(userID int, avatarFile io.Reader, avatarPath string) (string, error) {
	storageAvatarPath := UserAvatarDirectory + "/" + strconv.Itoa(userID) + filepath.Ext(avatarPath)
	storageAvatarFile, err := os.OpenFile(storageAvatarPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer storageAvatarFile.Close()
	io.Copy(storageAvatarFile, avatarFile)
	return storageAvatarPath, nil
}

func (u userUsecase) deleteFileIfExists(fileName string) {
	if u.fileExists(fileName) {
		os.Remove(fileName)
	}
}

func (u userUsecase) fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}