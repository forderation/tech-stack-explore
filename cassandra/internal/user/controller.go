package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/forderation/cassandra-learn/internal/pkg/cryptography"
	"github.com/forderation/cassandra-learn/internal/pkg/storage"
)

type Controller struct {
	UserManager  storage.UserManager
	Cryptography cryptography.Cryptography
}

func (c Controller) List(w http.ResponseWriter, r *http.Request) {
	var req Request
	req.bind(r.URL)

	var page []byte
	if req.PageCursor != "" {
		var err error
		page, err = c.Cryptography.DecryptString(req.PageCursor, nil)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid page cursor"))
			return
		}
	}

	users, page, err := c.UserManager.List(r.Context(), req.PageSize, page)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}

	var cursor string
	if len(page) != 0 {
		log.Println("page before encrypt:", string(page))
		cursor, err = c.Cryptography.EncryptAsString(page, nil)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal error"))
			return
		}
	}

	var res Response
	res.bind(users, len(users), cursor, req.PageSize)
	body, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)
}
