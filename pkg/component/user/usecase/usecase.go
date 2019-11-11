package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"io"
	"strings"
)

const PasswordSaltLength = 8

type userUsecase struct {
	repository        user.Repository
	fileRepository    user.FileRepository
	sessionRepository user.SessionRepository
}

func NewUserUsecase(
	repository user.Repository,
	fileRepository user.FileRepository,
	sessionRepository user.SessionRepository) user.Usecase {
	return &userUsecase{
		repository:        repository,
		fileRepository:    fileRepository,
		sessionRepository: sessionRepository,
	}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	return u.repository.GetAll()
}

func (u *userUsecase) GetUserBySessionID(sessionID string) (*model.User, error) {
	userID, ok := u.sessionRepository.GetUserBySession(sessionID)
	if !ok {
		return nil, nil
	}
	return u.repository.GetByID(userID)
}

func (u *userUsecase) UpdateUser(id int, password, name string) error {
	userToUpdate, err := u.repository.GetByID(id)
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
	return u.repository.Update(*userToUpdate)
}

func (u *userUsecase) UpdateUserAvatar(user *model.User, avatarFile io.Reader, avatarPath string) error {
	newAvatarPath, err := u.fileRepository.StoreAvatar(user, avatarFile, avatarPath)
	if err != nil {
		return err
	}
	return u.repository.UpdateAvatarPath(user.ID, newAvatarPath)
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
	userWithSameEmail, err := u.repository.GetByEmail(email)
	if err != nil {
		return errors.New("unknown error")
	}
	if userWithSameEmail != nil {
		return fmt.Errorf("user with email %v already registered", email)
	}
	passwordHash := u.getPasswordHash(password)
	return u.repository.Create(email, passwordHash, name)
}

func (u *userUsecase) Login(email, password string) (sessionID string, err error) {
	userToLogin, err := u.repository.GetByEmail(email)
	if userToLogin == nil {
		return "", fmt.Errorf("incorrect email")
	}
	if err != nil {
		return "", fmt.Errorf("unknown error")
	}
	if !u.isPasswordsEqual(password, userToLogin.PasswordHash) {
		return "", fmt.Errorf("incorrect password")
	}
	sessionID = u.getNewSessionID()
	err = u.sessionRepository.AddSession(sessionID, userToLogin.ID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

func (u *userUsecase) Logout(sessionID string) error {
	err := u.sessionRepository.DeleteSession(sessionID)
	if err != nil {
		return fmt.Errorf("error deleting session")
	}
	return nil
}

func (u *userUsecase) getPasswordHash(password string) string {
	salt := make([]byte, PasswordSaltLength)
	rand.Read(salt)
	return u.getPasswordHashWithSalt(password, salt)
}

func (u *userUsecase) isPasswordsEqual(password string, passwordHash string) bool {
	decodedPasswordHash, _ := base64.RawStdEncoding.DecodeString(passwordHash)
	return u.getPasswordHashWithSalt(password, decodedPasswordHash[0:PasswordSaltLength]) == passwordHash
}

func (u *userUsecase) getPasswordHashWithSalt(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func (u *userUsecase) getNewSessionID() string {
	return uuid.New().String()
}
