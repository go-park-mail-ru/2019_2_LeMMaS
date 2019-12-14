package grpc

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
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
	auth.RegisterUserServer(h.server, &h)
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
