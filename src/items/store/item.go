package store

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/items/model"
)

type ItemStore interface {
	GetItem(ctx context.Context, itemId int64) (*model.Item, error)
}
