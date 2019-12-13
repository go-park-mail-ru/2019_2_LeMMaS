package repository

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/model"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
)

type userRepository struct {
}

func NewUserRepository() auth.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(email string, passwordHash string, name string) error {
	// todo: call user microservice
	return nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	// todo: call user microservice
	return nil, nil
}
