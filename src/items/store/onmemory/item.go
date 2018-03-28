package onmemory

import (
	"context"
	"fmt"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/items/model"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/items/store"
)

type onmemoryItemStore struct{}

var items = map[int64]*model.Item{
	1: &model.Item{
		Id:   1,
		Name: "RealForce",
	},
	2: &model.Item{
		Id:   2,
		Name: "Majestouch",
	},
	3: &model.Item{
		Id:   3,
		Name: "Happy Hacking Keyboard",
	},
}

func NewItemStore() store.ItemStore {
	return &onmemoryItemStore{}
}

func (s *onmemoryItemStore) GetItem(ctx context.Context, itemId int64) (*model.Item, error) {
	item, ok := items[itemId]
	if !ok {
		return nil, fmt.Errorf("item_id %d is not found", itemId)
	}
	return item, nil
}
