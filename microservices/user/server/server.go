package main

import (
	"fmt"
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/user/proto"
	user "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/user/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("cant listen port", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServer(server, user.NewUserManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}