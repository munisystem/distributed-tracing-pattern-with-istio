package onmemory

import (
	"context"
	"fmt"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/model"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/store"
)

type onmemoryReviewStore struct{}

var reviews = map[int64]*model.Review{
	1: &model.Review{
		Id:          1,
		CustomerId:  1,
		ItemId:      2,
		Rating:      4,
		Description: "I bought cherry mx blue. I'm having good experience of typing, but noisy.",
	},
	2: &model.Review{
		Id:          2,
		CustomerId:  2,
		ItemId:      3,
		Rating:      5,
		Description: "HHKB is compact compared with other keyboards and good design. I love it.",
	},
	3: &model.Review{
		Id:          3,
		CustomerId:  3,
		ItemId:      1,
		Rating:      4,
		Description: "This is good product. It's only fault that size is bigger.",
	},
}

func NewReviewStore() store.ReviewStore {
	return &onmemoryReviewStore{}
}

func (s *onmemoryReviewStore) GetCustomerReview(ctx context.Context, customerId int64) (*model.Review, error) {
	r, ok := reviews[customerId]
	if !ok {
		return nil, fmt.Errorf("customer_id %d is not reviewed.", customerId)
	}
	return r, nil
}
