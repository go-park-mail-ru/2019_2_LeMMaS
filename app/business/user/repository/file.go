package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/app/model"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

const (
	UserAvatarDirectory = "static/user/avatar"

	FilePerm      = 0666
	DirectoryPerm = 0777
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) StoreAvatar(user *model.User, avatarFile io.Reader, avatarPath string) (string, error) {
	if err := os.MkdirAll(r.getPath(UserAvatarDirectory), DirectoryPerm); err != nil {
		logger.Error(err)
		return "", err
	}
	if err := r.deleteFileIfExists(user.AvatarPath); err != nil {
		logger.Error(err)
		return "", err
	}
	storageAvatarPath := uuid.New().String() + filepath.Ext(avatarPath)
	fullStorageAvatarPath := r.getPath(UserAvatarDirectory) + "/" + storageAvatarPath
	storageAvatarFile, err := os.OpenFile(fullStorageAvatarPath, os.O_WRONLY|os.O_CREATE, FilePerm)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	defer storageAvatarFile.Close()
	io.Copy(storageAvatarFile, avatarFile)
	return storageAvatarPath, nil
}

func (r *fileRepository) getPath(directory string) string {
	serverRoot := os.Getenv("SERVER_ROOT")
	if serverRoot == "" {
		return directory
	}
	return serverRoot + "/" + directory
}

func (r *fileRepository) deleteFileIfExists(fileName string) error {
	if r.fileExists(fileName) {
		return os.Remove(fileName)
	}
	return nil
}

func (r *fileRepository) fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
