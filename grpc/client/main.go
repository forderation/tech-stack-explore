package client

import (
	"context"

	"github.com/forderation/tech-stack-explore/grpc/common/model"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type UsersClient interface {
	Register(ctx context.Context, in *model.User, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*model.UserList, error)
}

func main() {

}
