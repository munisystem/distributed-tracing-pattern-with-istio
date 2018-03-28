package main

import (
	"os"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/client/items"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/server"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/store/onmemory"
	"go.uber.org/zap"
)

func main() {
	itemsUrl := os.Getenv("ITEMS_URL")
	if itemsUrl == "" {
		itemsUrl = "http://127.0.0.1:8083"
	}
	ic := items.New(itemsUrl)
	rs := onmemory.NewReviewStore()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	srv := server.New(ic, rs, logger)
	panic(srv.Run(":8082"))
}
