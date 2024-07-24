package proxy

import "proxy/internal/domain/proxy"

type Configuration func(s *Service) error

type Service struct {
	proxyRepository proxy.Repository
}

func New(configs ...Configuration) (s *Service, err error) {
	s = &Service{}

	for _, cfg := range configs {
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithProxyRepository(proxyRepository proxy.Repository) Configuration {
	return func(s *Service) error {
		s.proxyRepository = proxyRepository
		return nil
	}
}
