package store

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
)

type CustomerStore interface {
	GetCustomer(ctx context.Context, customerId int64) (*model.Customer, error)
}
