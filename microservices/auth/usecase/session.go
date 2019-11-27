package usecase

import (
	"encoding/base64"
	"fmt"
	user "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth"
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	repo "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/repository"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
	"golang.org/x/net/context"
	"math/rand"
	"os"
)

const passwordSaltLength = 8

type AuthManager struct {
	repository        user.Repository
	sessionRepository user.SessionRepository
}

func NewAuthManager() *AuthManager {
	conn, err := newRedis()
	if err != nil {
		log.Error(err)
	}
	sessionRepository := repo.NewSessionRepository(conn)

	db, err := newDB()
	if err != nil {
		log.Error(err)
	}
	repository := repo.NewDatabaseRepository(db)

	return &AuthManager{
		repository:        repository,
		sessionRepository: sessionRepository,
	}
}

func (sm *AuthManager) Login(ctx context.Context, userAuth *pb.UserAuth) (*pb.SessionIDAndError, error) {
	userToLogin, err := sm.repository.GetByEmail(userAuth.Email)
	if userToLogin == nil {
		errMes := "incorrect email"
		return &pb.SessionIDAndError{&pb.SessionID{""}, &pb.Error{errMes}}, fmt.Errorf(errMes)
	}
	if err != nil {
		errMes := "incorrect email"
		return &pb.SessionIDAndError{&pb.SessionID{""}, &pb.Error{errMes}}, fmt.Errorf(errMes)
	}
	if !sm.isPasswordsEqual(userAuth.Password, userToLogin.PasswordHash) {
		errMes := "incorrect email"
		return &pb.SessionIDAndError{&pb.SessionID{""}, &pb.Error{errMes}}, fmt.Errorf(errMes)
	}
	sessionID := sm.getNewSessionID()
	err = sm.sessionRepository.AddSession(sessionID, userToLogin.ID)
	if err != nil {
		errMes := "incorrect email"
		return &pb.SessionIDAndError{&pb.SessionID{""}, &pb.Error{errMes}}, fmt.Errorf(errMes)
	}
	return &pb.SessionIDAndError{&pb.SessionID{sessionID}, &pb.Error{"ok"}}, nil
}

func (sm *AuthManager) Logout(ctx context.Context, sessionID *pb.SessionID) (*pb.Error, error) {
	err := sm.sessionRepository.DeleteSession(sessionID.ID)
	if err != nil {
		return &pb.Error{""}, errors.New("error deleting session")
	}
	return &pb.Error{""}, nil
}

func (sm *AuthManager) RegisterUser(ctx context.Context, userDataRegister *pb.UserDataRegister) (*pb.Error, error) {
	fmt.Printf("register ready")
	userWithSameEmail, err := sm.repository.GetByEmail(userDataRegister.Email)
	if err != nil {
		errMes := "unknown error"
		return &pb.Error{errMes}, errors.New(errMes)
	}
	if userWithSameEmail != nil {
		errMes := "user with this email already registered"
		return &pb.Error{errMes}, errors.New(errMes)
	}
	passwordHash := sm.getPasswordHash(userDataRegister.Password)

	err = sm.repository.Create(userDataRegister.Email, passwordHash, userDataRegister.Name)
	if err != nil {
		return &pb.Error{"unknown error"}, err
	}
	return &pb.Error{"ok"}, nil
}

func (sm *AuthManager) getPasswordHash(password string) string {
	salt := make([]byte, passwordSaltLength)
	rand.Read(salt)
	return sm.getPasswordHashWithSalt(password, salt)
}

func (sm *AuthManager) isPasswordsEqual(password string, passwordHash string) bool {
	decodedPasswordHash, _ := base64.RawStdEncoding.DecodeString(passwordHash)
	return sm.getPasswordHashWithSalt(password, decodedPasswordHash[0:passwordSaltLength]) == passwordHash
}

func (sm *AuthManager) getPasswordHashWithSalt(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	hash = append(salt, hash...)
	return base64.RawStdEncoding.EncodeToString(hash)
}

func (sm *AuthManager) getNewSessionID() string {
	return uuid.New().String()
}

func newRedis() (redis.Conn, error) {
	key := os.Getenv("REDIS_DSN")
	if key == "" {
		key = "redis://redis:6379"
	}
	connection, err := redis.DialURL(key)
	if err != nil {
		return nil, err
	}
	_, err = connection.Do("PING")
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func newDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
