package main

import (
	"fmt"
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	auth "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

 // TODO for client in main.go
	//grcpConn, err := grpc.Dial(
	//		"127.0.0.1:8081",
	//		grpc.WithInsecure(),
	//	)
	//if err != nil {
	//	log.Fatalf("cant connect to grpc")
	//}
	//defer grcpConn.Close()

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("cant listen port", err)
	}

	server := grpc.NewServer()

	pb.RegisterAuthServer(server, auth.NewAuthManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
