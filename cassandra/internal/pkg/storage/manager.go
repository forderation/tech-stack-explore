package storage

import "context"

type UserManager interface {
	List(ctx context.Context, size int, page []byte) ([]*User, []byte, error)
}
