package storage

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

const UserAvatarDirectory = "_storage/user/avatar"

type FileStorage struct {
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (s FileStorage) StoreUserAvatar(userID int, avatarFile io.Reader, avatarPath string) (string, error) {
	storageAvatarPath := UserAvatarDirectory + "/" + strconv.Itoa(userID) + filepath.Ext(avatarPath)
	storageAvatarFile, err := os.OpenFile(storageAvatarPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer storageAvatarFile.Close()
	io.Copy(storageAvatarFile, avatarFile)
	return storageAvatarPath, nil
}

func (s FileStorage) DeleteFileIfExists(fileName string) {
	if s.fileExists(fileName) {
		os.Remove(fileName)
	}
}

func (s FileStorage) fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
