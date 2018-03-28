package client

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
)

type Reviews interface {
	GetCustomerReview(ctx context.Context, customerId int64) (*model.Review, error)
}
