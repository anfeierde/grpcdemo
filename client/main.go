package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	auth "thebox/pb"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := auth.NewAuthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := client.Login(ctx, &auth.LoginRequest{
		Username: "wk",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	log.Println(resp.Token)
}
