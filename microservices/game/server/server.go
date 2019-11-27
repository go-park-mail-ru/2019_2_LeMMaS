package main

import (
	"fmt"
	pb "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/proto"
	game "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/game/usecase"
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

	pb.RegisterGameServer(server, game.NewGameManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
