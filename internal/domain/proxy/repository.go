package proxy

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, id string) (res Entity, err error)
	Add(ctx context.Context, data Entity) (dest string, err error)
}
