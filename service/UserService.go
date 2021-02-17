package main

import (
	"log"
	"mygo/gen/proto/user"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	PORT = ":9001"
)

type server struct{}

func (s *server) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	return &user.LoginResp{Result: "hello, " + req.Username}, nil
}

func main() {
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
