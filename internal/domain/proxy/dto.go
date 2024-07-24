package proxy

import (
	"errors"
	"net/http"
)

type Request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Method == "" {
		return errors.New("method: cannot be blank")
	}

	if s.Url == "" {
		return errors.New("url: cannot be blank")
	}

	if s.Headers == nil {
		return errors.New("headers: cannot be blank")
	}

	return nil
}

type Response struct {
	ID      string              `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int64               `json:"length"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:      data.ID,
		Status:  data.Response.Status,
		Headers: data.Response.Headers,
		Length:  data.Response.Length,
	}
	return
}
