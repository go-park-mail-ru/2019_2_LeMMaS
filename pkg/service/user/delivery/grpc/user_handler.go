package grpc

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type UserHandler struct {
	server *grpc.Server
}

func NewUserHandler() *UserHandler {
	h := UserHandler{}
	h.server = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	user.RegisterUserServer(h.server, &h)
	return &h
}

func (h *UserHandler) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Println("listening " + address)
	return h.server.Serve(listener)
}

func (h *UserHandler) GetByID(context.Context, *user.GetByIDParams) (*user.GetByIDResult, error) {

}

func (h *UserHandler) GetBySession(context.Context, *user.GetBySessionParams) (*user.GetBySessionResult, error) {

}

func (h *UserHandler) Update(context.Context, *user.UpdateParams) (*user.UpdateResult, error) {

}

func (h *UserHandler) UpdateAvatar(context.Context, *user.UpdateAvatarParams) (*user.UpdateAvatarResult, error) {

}

func (h *UserHandler) GetSpecialAvatar(context.Context, *user.GetSpecialAvatarParams) (*user.GetSpecialAvatarResult, error) {

}
