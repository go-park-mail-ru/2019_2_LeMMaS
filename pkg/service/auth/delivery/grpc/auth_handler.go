package grpc

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"os"
	"time"
)

type AuthHandler struct {
	server *grpc.Server
}

func NewAuthHandler(s auth.AuthServer) *AuthHandler {
	h := AuthHandler{}
	h.server = grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	auth.RegisterAuthServer(h.server, s)
	return &h
}

func (h *AuthHandler) Serve() error {
	listener, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("cant listen port: %w", err)
	}
	return h.server.Serve(listener)
}
