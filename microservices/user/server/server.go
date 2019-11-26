package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("cant listet port", err)
	}

	server := grpc.NewServer()

	proto.RegisterAuthServer(server, pb.NewUserManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}