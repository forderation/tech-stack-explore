package cassandra

import (
	"context"
	"time"

	"github.com/forderation/cassandra-learn/internal/pkg/storage"
	"github.com/gocql/gocql"
)

var _ storage.UserManager = User{}

type User struct {
	Connection *gocql.Session
	Timeout    time.Duration
}

func (u User) List(ctx context.Context, size int, page []byte) ([]*storage.User, []byte, error) {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()
	query := `SELECT id, username, created_at FROM users`
	itr := u.Connection.Query(query).WithContext(ctx).PageSize(size).PageState(page).Iter()
	defer itr.Close()

	page = itr.PageState()
	users := make([]*storage.User, 0, itr.NumRows())
	scanner := itr.Scanner()
	for scanner.Next() {
		user := &storage.User{}
		if err := scanner.Scan(
			&user.ID,
			&user.Username,
			&user.CreatedAt,
		); err != nil {
			return nil, nil, err
		}
		users = append(users, user)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return users, page, nil
}
