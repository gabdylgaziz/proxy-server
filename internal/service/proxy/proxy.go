package proxy

import (
	"context"
	"net/http"
	"proxy/internal/domain/proxy"
)

func (s *Service) GetResponse(ctx context.Context, id string) (res proxy.Response, err error) {
	data, err := s.proxyRepository.Get(ctx, id)
	if err != nil {
		return
	}

	res = proxy.ParseFromEntity(data)
	return
}

func (s *Service) CreateRequest(ctx context.Context, req proxy.Request) (res proxy.Response, err error) {
	data := proxy.Entity{
		Request: req,
	}

	request, err := http.NewRequest(data.Request.Method, data.Request.Url, nil)
	if err != nil {
		return
	}

	for key, value := range data.Request.Headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data.Response = proxy.Response{
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Length:  resp.ContentLength,
	}

	id, err := s.proxyRepository.Add(ctx, data)
	data.ID = id

	res = proxy.ParseFromEntity(data)

	return
}
