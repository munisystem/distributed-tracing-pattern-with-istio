package client

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/model"
)

type Items interface {
	GetItem(ctx context.Context, itemId int64) (*model.Item, error)
}
