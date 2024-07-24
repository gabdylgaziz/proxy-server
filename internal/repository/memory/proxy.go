package memory

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"proxy/internal/domain/proxy"
	"sync"
)

type ProxyRepository struct {
	db sync.Map
}

func NewProxyRepository() *ProxyRepository {
	return &ProxyRepository{}
}

func (r *ProxyRepository) Get(ctx context.Context, id string) (res proxy.Entity, err error) {
	resp, ok := r.db.Load(id)
	if !ok {
		err = sql.ErrNoRows
		return
	}
	res = resp.(proxy.Entity)

	return
}

func (r *ProxyRepository) Add(ctx context.Context, data proxy.Entity) (dest string, err error) {
	requestID := uuid.New().String()
	data.ID = requestID
	r.db.Store(data.ID, data)

	return data.ID, nil
}
