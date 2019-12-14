package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	repository2 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth/repository"
	user3 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/proto"
	repository3 "github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/repository"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/argon2"
	"golang.org/x/net/context"
	"os"
	"os/user"
	"strings"
)

const passwordSaltLength = 8

type UserManager struct {
	repository        user3.Repository
	fileRepository    user3.FileRepository
	sessionRepository user3.SessionRepository
}

func (u *UserManager) GetUserByID(ctx context.Context, userID *user.UserID) (*user.UserAndError, error) {
	user, err := u.repository.GetByID(int(userID.ID))
	if err != nil {
		return &user.UserAndError{}, err
	}
	pbUser := &user.UserData{user.Email, user.PasswordHash, user.Name}
	return &user.UserAndError{pbUser, &user.Error{"ok"}}, err
}

func (u *UserManager) GetUserBySessionID(ctx context.Context, sessionID *user.SessionID) (*user.UserAndError, error) {
	userID, ok := u.sessionRepository.GetUserBySession(sessionID.Id)
	if !ok {
		return nil, nil
	}
	return u.GetUserByID(ctx, &user.UserID{int32(userID)})
}

func (u *UserManager) UpdateUser(ctx context.Context, pbUserToUpdate *user.UserToUpdate) (*user.Error, error) {
	userToUpdate, err := u.repository.GetByID(int(pbUserToUpdate.UserID))
	if err != nil {
		return &user.Error{"unknown error"}, err
	}
	if pbUserToUpdate.Password != "" {
		userToUpdate.PasswordHash = u.getPasswordHash(pbUserToUpdate.Password)
	}
	if pbUserToUpdate.Name != "" && userToUpdate.Name != pbUserToUpdate.Name {
		userToUpdate.Name = pbUserToUpdate.Name
		avatarPath, _ := u.GetAvatarUrlByName(ctx, &user.UserName{pbUserToUpdate.Name})
		if avatarPath.AvatarUrl != "" {
			userToUpdate.AvatarPath = avatarPath.AvatarUrl
		}
	}
	err = u.repository.Update(*userToUpdate)
	if err != nil {
		return &user.Error{"unknown error"}, err
	}
	return &user.Error{"ok"}, nil
}

func (u *UserManager) UpdateUserAvatar(ctx context.Context, user *user.UserToUpdateAvatar) (*user.Error, error) {
	// TODO преобразование типов
	newAvatarPath, err := u.fileRepository.Store(user.AvatarFile)
	if err != nil {
		return &user.Error{"unknown error"}, err
	}
	err = u.repository.UpdateAvatarPath(int(user.UserID), newAvatarPath)
	if err != nil {
		return &user.Error{"unknown error"}, err
	}
	return &user.Error{"ok"}, nil
}

func (u *UserManager) GetLeaderUsers(ctx context.Context, userID *user.UserID) (*user.Users, error) {
	users, err := u.repository.GetAll(int(userID.ID))
	pbUsers := translateType(users)
	return pbUsers, err
}

func translateType(users []model.User) *user.Users {
	pbUsers := &user.Users{}
	for _, user := range users {
		pbUsers = append(pbUsers, &user.UserData{user.Email, user.PasswordHash, user.Name})
	}
	return pbUsers
}

func (u *UserManager) GetAvatarUrlByName(ctx context.Context, userName *user.UserName) (*user.AvatarUrl, error) {
	avatarsByName := map[string]string{
		"earth":   "http://www.i2clipart.com/cliparts/3/d/1/e/clipart-earth-3d1e.png",
		"trump":   "https://lemmas.s3.eu-west-3.amazonaws.com/trump.png",
		"lebedev": "https://lemmas.s3.eu-west-3.amazonaws.com/lebedev.jpg",
		"cat":     "https://i.pinimg.com/originals/90/a8/56/90a856d434dd9df24d8d5fdf4bf3ce72.png",
	}
	return &user.AvatarUrl{avatarsByName[strings.ToLower(userName.Name)]}, nil
}

func (u *UserManager) getPasswordHash(password string) string {
	salt := make([]byte, passwordSaltLength)
	rand.Read(salt)
	return u.getPasswordHashWithSalt(password, salt)
}

func (u *UserManager) getPasswordHashWithSalt(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func newRedis() (redis.Conn, error) {
	key := os.Getenv("REDIS_DSN")
	fmt.Println(key)
	if key == "" {
		key = "redis://redis:6379"
	}
	connection, err := redis.DialURL(key)
	if err != nil {
		return nil, err
	}
	_, err = connection.Do("PING")
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func newDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
