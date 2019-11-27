package main

import (
	"fmt"
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	auth "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/usecase"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("cant listen port", err)
	}

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))

	pb.RegisterAuthServer(server, auth.NewAuthManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
