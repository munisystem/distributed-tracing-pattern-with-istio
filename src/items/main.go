package main

import (
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/items/server"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/items/store/onmemory"
)

func main() {
	is := onmemory.NewItemStore()
	svr := server.New(is)
	panic(svr.Run(":8083"))
}
