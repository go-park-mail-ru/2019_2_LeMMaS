package usecase

import (
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"golang.org/x/net/context"
	"sync"
)

// TODO realisations of server handlers

type SessionManager struct {
	mu       sync.RWMutex
	sessions map[pb.SessionID]*pb.Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu:       sync.RWMutex{},
		sessions: map[pb.SessionID]*pb.Session{},
	}
}

func (sm *SessionManager) CreateSession(ctx context.Context, session *pb.Session) (pb.SessionID, error) {

}

func (sm *SessionManager) CheckSession(ctx context.Context, session *pb.SessionID) (pb.Session, error) {

}

func (sm *SessionManager) DeleteSession(ctx context.Context, session *pb.SessionID) (pb.Error, error) {

}
