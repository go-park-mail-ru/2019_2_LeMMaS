package auth

import (
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

type Service interface {
	Login(ctx context.Context, session *pb.UserAuth) (pb.SessionIDAndError, error)
	Logout(ctx context.Context, session *pb.SessionID) (pb.Error, error)
	CheckSession(ctx echo.Context, session *pb.SessionID) (pb.Session, error)
	RegisterUser(ctx echo.Context, session *pb.UserDataRegister) (pb.Error, error)
}
