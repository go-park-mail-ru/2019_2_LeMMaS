package usecase

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"golang.org/x/net/context"
	"sync"
)

// TODO realisations of server handlers

type SessionManager struct {
	mu       sync.RWMutex
	sessions map[proto.SessionID]*proto.Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu:       sync.RWMutex{},
		sessions: map[proto.SessionID]*proto.Session{},
	}
}

func (sm *SessionManager) CreateSession(ctx context.Context, session *proto.Session) (proto.SessionID, error) {

}

func (sm *SessionManager) CheckSession(ctx context.Context, session *proto.SessionID) (proto.Session, error) {

}

func (sm *SessionManager) DeleteSession(ctx context.Context, session *proto.SessionID) (proto.Error, error) {

}