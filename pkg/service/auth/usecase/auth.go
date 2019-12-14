package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

const passwordSaltLength = 8

type authUsecase struct {
	userRepo    auth.UserRepository
	sessionRepo auth.SessionRepository
	logger      logger.Logger
}

func NewAuthUsecase(userRepo auth.UserRepository, sessionRepo auth.SessionRepository, logger logger.Logger) auth.AuthUsecase {
	return &authUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		logger:      logger,
	}
}

func (u *authUsecase) Login(email, password string) (session string, err error) {
	userToLogin, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return
	}
	if userToLogin == nil {
		err = errors.New("incorrect email")
		return
	}
	if !u.isPasswordsEqual(password, userToLogin.PasswordHash) {
		err = errors.New("incorrect password")
		return
	}
	session = u.newSessionID()
	err = u.sessionRepo.Add(session, userToLogin.ID)
	return
}

func (u *authUsecase) Logout(session string) error {
	return u.sessionRepo.Delete(session)
}

func (u *authUsecase) Register(email, password, name string) error {
	userWithSameEmail, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return err
	}
	if userWithSameEmail != nil {
		return errors.New("user with this email already registered")
	}
	return u.userRepo.Register(email, u.getPasswordHash(password), name)
}

func (u *authUsecase) GetUser(session string) (int, bool) {
	return u.sessionRepo.Get(session)
}

func (u *authUsecase) getPasswordHash(password string) string {
	salt := make([]byte, passwordSaltLength)
	rand.Read(salt)
	return u.getPasswordHashWithSalt(password, salt)
}

func (u *authUsecase) isPasswordsEqual(password string, passwordHash string) bool {
	decodedPasswordHash, _ := base64.RawStdEncoding.DecodeString(passwordHash)
	return u.getPasswordHashWithSalt(password, decodedPasswordHash[0:passwordSaltLength]) == passwordHash
}

func (u *authUsecase) getPasswordHashWithSalt(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func (u *authUsecase) newSessionID() string {
	return uuid.New().String()
}
