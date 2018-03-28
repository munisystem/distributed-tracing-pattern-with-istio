package client

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
)

type Histories interface {
	GetCustomerHistory(ctx context.Context, customerId int64) (*model.History, error)
}
