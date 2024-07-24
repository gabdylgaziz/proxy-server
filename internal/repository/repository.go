package repository

import (
	"proxy/internal/domain/proxy"
	"proxy/internal/repository/memory"
)

type Configuration func(r *Repository) error

type Repository struct {
	Proxy proxy.Repository
}

func New(configs ...Configuration) (s *Repository, err error) {
	s = &Repository{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithMemoryStore() Configuration {
	return func(s *Repository) (err error) {
		s.Proxy = memory.NewProxyRepository()
		return
	}
}
