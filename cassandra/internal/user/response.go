package user

import "fmt"

type Response struct {
	Data interface{} `json:"data"`
	Meta struct {
		Total int `json:"total,omitempty"`
	} `json:"meta"`
	Links struct {
		Next string `json:"next,omitempty"`
	} `json:"links"`
}

func (r *Response) bind(data interface{}, total int, cursor string, size int) {
	r.Data = data
	r.Meta.Total = total

	if cursor != "" {
		r.Links.Next += fmt.Sprintf("page[cursor]=%s&", cursor)
	}
	if size != 0 {
		r.Links.Next += fmt.Sprintf("page[size]=%d&", size)
	}
	if r.Links.Next != "" {
		r.Links.Next = "?" + r.Links.Next[:len(r.Links.Next)-1]
	}
}
