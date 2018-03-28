package store

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/model"
)

type ReviewStore interface {
	GetCustomerReview(ctx context.Context, customerId int64) (*model.Review, error)
}
