package component

import (
	"crypto/md5"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/storage"
	"github.com/google/uuid"
	"io"
)

type UserComponent struct {
	userStorage *storage.UserStorage
	fileStorage *storage.FileStorage
	sessions    map[string]int
}

func NewUserComponent() *UserComponent {
	return &UserComponent{
		userStorage: storage.NewUserStorage(),
		fileStorage: storage.NewFileStorage(),
		sessions:    map[string]int{},
	}
}

func (c UserComponent) GetAllUsers() []storage.User {
	return c.userStorage.GetAll()
}

func (c UserComponent) GetUserBySessionID(sessionID string) *storage.User {
	userID := c.sessions[sessionID]
	return c.userStorage.GetByID(userID)
}

func (c UserComponent) UpdateUser(id int, password, name string) {
	passwordHash := ""
	if password != "" {
		passwordHash = c.getPasswordHash(password)
	}
	c.userStorage.Update(id, passwordHash, name)
}

func (c UserComponent) UpdateUserAvatar(user *storage.User, avatarFile io.Reader, avatarPath string) error {
	c.fileStorage.DeleteFileIfExists(user.AvatarPath)
	newAvatarPath, err := c.fileStorage.StoreUserAvatar(user.ID, avatarFile, avatarPath)
	if err != nil {
		return err
	}
	c.userStorage.UpdateAvatarPath(user.ID, newAvatarPath)
	return nil
}

func (c UserComponent) Register(email, password, name string) error {
	if c.userStorage.GetByEmail(email) != nil {
		return fmt.Errorf("user with email %v already registered", email)
	}
	passwordHash := c.getPasswordHash(password)
	c.userStorage.Create(email, passwordHash, name)
	return nil
}

func (c UserComponent) Login(email, password string) (string, error) {
	user := c.userStorage.GetByEmail(email)
	if user == nil {
		return "", fmt.Errorf("incorrect email")
	}
	if c.getPasswordHash(password) != user.PasswordHash {
		return "", fmt.Errorf("incorrect password")
	}
	sessionID := c.getNewSessionID()
	c.sessions[sessionID] = user.ID
	return sessionID, nil
}

func (c UserComponent) Logout(sessionID string) error {
	if _, ok := c.sessions[sessionID]; !ok {
		return fmt.Errorf("session id not found")
	}
	delete(c.sessions, sessionID)
	return nil
}

func (c UserComponent) getPasswordHash(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func (c UserComponent) getNewSessionID() string {
	return uuid.New().String()
}
