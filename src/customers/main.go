package main

import (
	"os"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/client/histories"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/client/reviews"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/server"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/store/onmemory"
)

func main() {
	reviewsUrl := os.Getenv("REVIEWS_URL")
	if reviewsUrl == "" {
		reviewsUrl = "127.0.0.1:8082"
	}
	rc, err := reviews.New(reviewsUrl)
	if err != nil {
		panic(err)
	}

	historiesUrl := os.Getenv("HISTORIES_URL")
	if historiesUrl == "" {
		historiesUrl = "http://127.0.0.1:8081"
	}
	hc := histories.New(historiesUrl)

	cs := onmemory.NewCustomerStore()
	svr := server.New(hc, rc, cs)
	panic(svr.Run(":8080"))
}
