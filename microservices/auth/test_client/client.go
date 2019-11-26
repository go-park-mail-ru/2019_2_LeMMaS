package main

import (
	"fmt"
	session "github.com/go-park-mail-ru/2019_2_LeMMaS/microservices/auth/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	grcpConn, err := grpc.Dial(
			"127.0.0.1:8081",
			grpc.WithInsecure(),
		)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	sessManager := session.NewAuthClient(grcpConn)

	ctx := context.Background()

	userData := &session.UserDataRegister{
		"leda",
		"asdfgh",
		"leda",
	}
	_, err = sessManager.RegisterUser(ctx, userData)

	fmt.Println(err)
}
