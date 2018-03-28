package server

import (
	"context"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/items/store"
)

type Server struct {
	itemStore store.ItemStore
}

func New(itemStore store.ItemStore) *Server {
	return &Server{
		itemStore: itemStore,
	}
}

func (s *Server) Run(port string) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/api/items/:id", s.getItemHandler)
	return e.Start(port)
}

type getCustomerResponse struct {
	Message string `json:"message,omitempty"`
}

func (s *Server) getItemHandler(c echo.Context) error {
	itemId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(404, &getCustomerResponse{Message: "id is invalid"})
	}
	ctx := context.Background()
	item, err := s.itemStore.GetItem(ctx, itemId)
	if err != nil {
		return c.JSON(404, &getCustomerResponse{Message: "Not Found"})
	}
	return c.JSON(200, item)
}
