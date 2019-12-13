package grpc

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/auth"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type AuthHandler struct {
}

func NewAuthHandler(s auth.AuthServer) (*AuthHandler, error) {
	g := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))
	auth.RegisterAuthServer(g, s)
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		return nil, fmt.Errorf("cant listen port: %w", err)
	}
	err = g.Serve(listener)
	return &AuthHandler{}, err
}
