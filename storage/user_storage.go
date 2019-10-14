package storage

import (
	"sort"
	"sync"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
	Name         string
}

type UserStorage struct {
	usersByID map[int]User
	mutex     *sync.Mutex
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		usersByID: make(map[int]User),
		mutex:     &sync.Mutex{},
	}
}

func (s *UserStorage) Create(email string, passwordHash string, name string) {
	user := User{
		ID:           len(s.usersByID) + 1,
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
	}
	s.storeUser(&user)
}

func (s *UserStorage) Update(id int, passwordHash string, name string) {
	user := s.usersByID[id]
	if passwordHash != "" {
		user.PasswordHash = passwordHash
	}
	if name != "" {
		user.Name = name
	}
	s.usersByID[id] = user
}

func (s *UserStorage) GetAll() []User {
	result := make([]User, 0, len(s.usersByID))
	for _, user := range s.usersByID {
		result = append(result, user)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func (s *UserStorage) GetByEmail(email string) *User {
	for _, user := range s.usersByID {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func (s *UserStorage) storeUser(user *User) {
	s.mutex.Lock()
	s.usersByID[user.ID] = *user
	s.mutex.Unlock()
}
