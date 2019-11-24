package auth

import (
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"github.com/labstack/echo"
)

type Service interface {
	CreateSession(ctx echo.Context, session *proto.Session) (proto.SessionID, error)
	CheckSession(ctx echo.Context, session *proto.SessionID) (proto.Session, error)
	DeleteSession(ctx echo.Context, session *proto.SessionID) (proto.Error, error)
}
