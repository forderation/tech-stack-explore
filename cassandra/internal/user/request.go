package user

import (
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	PageCursor string
	PageSize   int
}

func (r *Request) bind(u *url.URL) {
	r.PageCursor = strings.ReplaceAll(u.Query().Get("page[cursor]"), " ", "")
	v, err := strconv.Atoi(u.Query().Get("page[size]"))
	switch {
	case err != nil, v < 1:
		r.PageSize = 10
	default:
		r.PageSize = v
	}
}
