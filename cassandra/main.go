package main

import (
	"log"
	"time"

	"github.com/forderation/cassandra-learn/internal/pkg/cryptography"
	"github.com/forderation/cassandra-learn/internal/pkg/http"
	"github.com/forderation/cassandra-learn/internal/pkg/storage/cassandra"
	"github.com/forderation/cassandra-learn/internal/user"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	con, err := cassandra.NewConnection(cassandra.ConnectionConfig{
		Hosts:        []string{"127.0.0.1"},
		Port:         9042,
		ProtoVersion: 4,
		Consistency:  "Quorum",
		Keyspace:     "blog",
		Timeout:      time.Second * 5,
	})

	if err != nil {
		log.Fatalln("connection error:", err)
	}

	defer con.Close()

	usr := cassandra.User{
		Connection: con,
		Timeout:    time.Second * 5,
	}

	cry := cryptography.New("b09e58536e4df2a4fc6dd3c9773e4f3d")

	ctr := user.Controller{
		UserManager:  usr,
		Cryptography: cry,
	}

	router := http.NewRouter()
	router.HandleFunc("/api/v1/users", ctr.List)

	srv := http.NewServer(router, ":8089")
	log.Println("start server on http://localhost:8089")
	log.Fatal(srv.ListenAndServe())
}
