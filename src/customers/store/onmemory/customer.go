package onmemory

import (
	"context"
	"fmt"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/store"
)

type onmemoryCustomerStore struct{}

var customers = map[int64]*model.Customer{
	1: &model.Customer{
		ID:    1,
		Name:  "alice",
		Email: "alice@example.com",
	},
	2: &model.Customer{
		ID:    2,
		Name:  "bob",
		Email: "bob@example.com",
	},
	3: &model.Customer{
		ID:    3,
		Name:  "carol",
		Email: "carol@example.com",
	},
}

func NewCustomerStore() store.CustomerStore {
	return &onmemoryCustomerStore{}
}

func (s *onmemoryCustomerStore) GetCustomer(ctx context.Context, customerId int64) (*model.Customer, error) {
	customer, ok := customers[customerId]
	if !ok {
		return nil, fmt.Errorf("customer_id %d is not found", customerId)
	}
	return customer, nil
}
