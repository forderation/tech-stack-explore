package main

import (
	"context"
	"log"
	"net"

	"github.com/forderation/tech-stack-explore/grpc/common/config"
	"github.com/forderation/tech-stack-explore/grpc/common/model"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct{}

func (UsersServer) Register(ctx context.Context, param *model.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Registering user", param.String())
	return new(empty.Empty), nil
}

func (UsersServer) List(ctx context.Context, void *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	grpcServer := grpc.NewServer()
	var userHandler UsersServer
	model.RegisterUsersServer(grpcServer, userHandler)
	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(grpcServer.Serve(listener))
}
