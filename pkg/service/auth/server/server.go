package server

import (
	"encoding/base64"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
	"golang.org/x/net/context"
	"math/rand"
)

const passwordSaltLength = 8

type authServer struct {
	userRepo    auth.UserRepository
	sessionRepo auth.SessionRepository
}

func NewAuthServer(userRepo auth.UserRepository, sessionRepo auth.SessionRepository) auth.AuthServer {
	return &authServer{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *authServer) Login(ctx context.Context, params *auth.LoginParams) (result *auth.LoginResult, err error) {
	userToLogin, err := s.userRepo.GetByEmail(params.Email)
	if err != nil {
		return
	}
	if userToLogin == nil {
		err = errors.New("incorrect email")
		return
	}
	if !s.isPasswordsEqual(params.Password, userToLogin.PasswordHash) {
		err = errors.New("incorrect password")
		return
	}
	sessionID := s.newSessionID()
	err = s.sessionRepo.Add(sessionID, userToLogin.ID)
	if err != nil {
		return
	}
	return &auth.LoginResult{SessionID: sessionID}, nil
}

func (s *authServer) Logout(ctx context.Context, params *auth.LogoutParams) (*auth.Nothing, error) {
	return nil, s.sessionRepo.Delete(params.SessionID)
}

func (s *authServer) RegisterUser(ctx context.Context, params *auth.RegisterParams) (result *auth.Nothing, err error) {
	userWithSameEmail, err := s.userRepo.GetByEmail(params.Email)
	if err != nil {
		return
	}
	if userWithSameEmail != nil {
		err = errors.New("user with this email already registered")
		return
	}
	err = s.userRepo.Create(params.Email, s.getPasswordHash(params.Password), params.Name)
	return
}

func (s *authServer) getPasswordHash(password string) string {
	salt := make([]byte, passwordSaltLength)
	rand.Read(salt)
	return s.getPasswordHashWithSalt(password, salt)
}

func (s *authServer) isPasswordsEqual(password string, passwordHash string) bool {
	decodedPasswordHash, _ := base64.RawStdEncoding.DecodeString(passwordHash)
	return s.getPasswordHashWithSalt(password, decodedPasswordHash[0:passwordSaltLength]) == passwordHash
}

func (s *authServer) getPasswordHashWithSalt(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func (s *authServer) newSessionID() string {
	return uuid.New().String()
}
