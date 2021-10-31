package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/forderation/tech-stack-explore/grpc/common/model"
)

type UsersServer interface {
	Register(context.Context, *model.User) (*empty.Empty, error)
	List(context.Context, *empty.Empty) (*model.UserList, error)
}

func main() {

}
