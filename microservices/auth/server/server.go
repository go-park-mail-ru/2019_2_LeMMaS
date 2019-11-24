package main

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	session "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

 // TODO for client in main.go
	//grcpConn, err := grpc.Dial(
	//	"127.0.0.1:8081",
	//	grpc.WithInsecure(),
	//)
	//if err != nil {
	//	log.Fatalf("cant connect to grpc")
	//	return router, err
	//}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("cant listet port", err)
	}

	server := grpc.NewServer()

	proto.RegisterAuthServer(server, session.NewSessionManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
