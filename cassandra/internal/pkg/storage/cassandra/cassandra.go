package cassandra

import (
	"time"

	"github.com/gocql/gocql"
)

type ConnectionConfig struct {
	Hosts        []string
	Port         int
	ProtoVersion int
	Consistency  string
	Keyspace     string
	Timeout      time.Duration
}

func NewConnection(config ConnectionConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.Hosts...)

	cluster.Port = config.Port
	cluster.ProtoVersion = config.ProtoVersion
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.ParseConsistency(config.Consistency)
	cluster.Timeout = config.Timeout

	return cluster.CreateSession()
}
