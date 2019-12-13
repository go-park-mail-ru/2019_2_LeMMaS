package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/proto"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/service/user/usecase"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("cant listen port", err)
	}

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute,
	}))

	user.RegisterUserServer(server, usecase.NewUserManager())

	fmt.Println("starting delivery at :8081")
	server.Serve(lis)
}
