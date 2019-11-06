package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

const (
	UserAvatarDirectory = "user/avatar"

	FilePerm      = 0666
	DirectoryPerm = 0777
)

type userFileRepository struct {
	staticPath string
}

func NewUserFileRepository() *userFileRepository {
	staticPath := os.Getenv("STATIC_PATH")
	if staticPath == "" {
		staticPath = "static"
	}
	return &userFileRepository{
		staticPath: staticPath,
	}
}

func (r *userFileRepository) StoreAvatar(user *model.User, avatarFile io.Reader, avatarPath string) (string, error) {
	if err := os.MkdirAll(r.getPath(UserAvatarDirectory), DirectoryPerm); err != nil {
		logger.Error(err)
		return "", err
	}
	if err := r.deleteFileIfExists(user.AvatarPath); err != nil {
		logger.Error(err)
		return "", err
	}
	storageAvatarPath := r.getPath(UserAvatarDirectory) + "/" + uuid.New().String() + filepath.Ext(avatarPath)
	storageAvatarFile, err := os.OpenFile(storageAvatarPath, os.O_WRONLY|os.O_CREATE, FilePerm)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	defer storageAvatarFile.Close()
	io.Copy(storageAvatarFile, avatarFile)
	return storageAvatarPath, nil
}

func (r *userFileRepository) getPath(directory string) string {
	return r.staticPath + "/" + directory
}

func (r *userFileRepository) deleteFileIfExists(fileName string) error {
	if r.fileExists(fileName) {
		return os.Remove(fileName)
	}
	return nil
}

func (r *userFileRepository) fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
