package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"google.golang.org/grpc"
	"net"
	"thebox/pb"
)

func main() {
	server := grpc.NewServer()
	auth.RegisterAuthServer(server, &User{})

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}

type User struct{}

func (u *User) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginResponse, error) {
	if r.Username == "wk" && r.Password == "123456" {
		h := md5.New()
		token := hex.EncodeToString(h.Sum([]byte(r.Username)))
		return &auth.LoginResponse{
			Token: token,
		}, nil
	}
	return nil, errors.New("username or password error")
}
