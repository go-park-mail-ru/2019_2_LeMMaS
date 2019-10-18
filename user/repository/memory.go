package repository

import (
	"sort"
	"sync"

	"github.com/go-park-mail-ru/2019_2_LeMMaS/model"
)

type memoryUserRepository struct {
	usersByID map[int]model.User
	mutex     *sync.Mutex
}

func NewMemoryUserRepository() *memoryUserRepository {
	return &memoryUserRepository{
		usersByID: make(map[int]model.User),
		mutex:     &sync.Mutex{},
	}
}

func (r *memoryUserRepository) Create(email string, passwordHash string, name string) {
	user := model.User{
		ID:           len(r.usersByID) + 1,
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
	}
	r.storeUser(&user)
}

func (r *memoryUserRepository) Update(id int, passwordHash string, name string) {
	user := r.usersByID[id]
	if passwordHash != "" {
		user.PasswordHash = passwordHash
	}
	if name != "" {
		user.Name = name
	}
	r.storeUser(&user)
}

func (r *memoryUserRepository) UpdateAvatarPath(id int, avatarPath string) {
	user := r.usersByID[id]
	user.AvatarPath = avatarPath
	r.storeUser(&user)
}

func (r *memoryUserRepository) GetAll() []model.User {
	result := make([]model.User, 0, len(r.usersByID))
	for _, user := range r.usersByID {
		result = append(result, user)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result
}

func (r *memoryUserRepository) GetByID(id int) *model.User {
	if user, ok := r.usersByID[id]; ok {
		return &user
	}
	return nil
}

func (r *memoryUserRepository) GetByEmail(email string) *model.User {
	for _, user := range r.usersByID {
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func (r *memoryUserRepository) storeUser(user *model.User) {
	r.mutex.Lock()
	r.usersByID[user.ID] = *user
	r.mutex.Unlock()
}
