package component

import (
	"crypto/md5"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/storage"
	"github.com/google/uuid"
)

type UserComponent struct {
	userStorage *storage.UserStorage
	sessions    map[string]*storage.User
}

func NewUserComponent() *UserComponent {
	return &UserComponent{
		userStorage: storage.NewUserStorage(),
		sessions:    map[string]*storage.User{},
	}
}

func (c UserComponent) GetAllUsers() []storage.User {
	return c.userStorage.GetAll()
}

func (c UserComponent) GetUserBySessionID(sessionID string) *storage.User {
	return c.sessions[sessionID]
}

func (c UserComponent) UpdateUser(id int, password, name string) {
	passwordHash := ""
	if password != "" {
		passwordHash = c.getPasswordHash(password)
	}
	c.userStorage.Update(id, passwordHash, name)
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
	c.sessions[sessionID] = user
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
