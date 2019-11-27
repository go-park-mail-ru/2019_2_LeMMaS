package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	authProto "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"io"
	"strings"
)

const passwordSaltLength = 8

type userUsecase struct {
	repository        user.Repository
	fileRepository    user.FileRepository
	sessionRepository user.SessionRepository
	auth              authProto.AuthClient
}

func NewUserUsecase(repository user.Repository, fileRepository user.FileRepository, sessionRepository user.SessionRepository, auth authProto.AuthClient) user.Usecase {
	return &userUsecase{
		repository:        repository,
		fileRepository:    fileRepository,
		sessionRepository: sessionRepository,
		auth:              auth,
	}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	return u.repository.GetAll()
}

func (u *userUsecase) GetUserByID(userID int) (*model.User, error) {
	return u.repository.GetByID(userID)
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

func (u *userUsecase) UpdateUserAvatar(user *model.User, avatarFile io.Reader) error {
	newAvatarPath, err := u.fileRepository.Store(avatarFile)
	if err != nil {
		return err
	}
	return u.repository.UpdateAvatarPath(user.ID, newAvatarPath)
}

func (u *userUsecase) GetAvatarUrlByName(name string) string {
	avatarsByName := map[string]string{
		"eath":    "http://www.i2clipart.com/cliparts/3/d/1/e/clipart-earth-3d1e.png",
		"trump":   "https://lemmas.s3.eu-west-3.amazonaws.com/trump.png",
		"lebedev": "https://lemmas.s3.eu-west-3.amazonaws.com/lebedev.jpg",
		"cat":     "https://i.pinimg.com/originals/90/a8/56/90a856d434dd9df24d8d5fdf4bf3ce72.png",
	}
	return avatarsByName[strings.ToLower(name)]
}

func (u *userUsecase) Register(email, password, name string) error {
	userData := &authProto.UserDataRegister{email, password, name}
	_, err := u.auth.RegisterUser(context.Background(), userData)
	return err
}

func (u *userUsecase) Login(email, password string) (sessionID string, err error) {
	userData := &authProto.UserAuth{email, password}
	result, err := u.auth.Login(context.Background(), userData)
	if err != nil {
		return "", err
	}
	return result.SessionID.ID, err
}

func (u *userUsecase) Logout(sessionID string) error {
	userData := &authProto.SessionID{sessionID}
	_, err := u.auth.Logout(context.Background(), userData)
	return err
}

func (u *userUsecase) getPasswordHash(password string) string {
	salt := make([]byte, passwordSaltLength)
	rand.Read(salt)
	return u.getPasswordHashWithSalt(password, salt)
}

func (u *userUsecase) isPasswordsEqual(password string, passwordHash string) bool {
	decodedPasswordHash, _ := base64.RawStdEncoding.DecodeString(passwordHash)
	return u.getPasswordHashWithSalt(password, decodedPasswordHash[0:passwordSaltLength]) == passwordHash
}

func (u *userUsecase) getPasswordHashWithSalt(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func (u *userUsecase) getNewSessionID() string {
	return uuid.New().String()
}
