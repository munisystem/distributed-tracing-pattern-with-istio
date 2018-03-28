package server

import (
	"context"
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/client"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/store"
)

type Server struct {
	historiesClient client.Histories
	reviewsClient   client.Reviews
	customerStore   store.CustomerStore
}

func New(historiesClient client.Histories, reviewsClient client.Reviews, customerStore store.CustomerStore) *Server {
	return &Server{
		historiesClient: historiesClient,
		reviewsClient:   reviewsClient,
		customerStore:   customerStore,
	}
}

func (s *Server) Run(port string) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/api/customers/:id", s.getCustomerhandler)
	return e.Start(port)
}

type getCustomerResponse struct {
	Message string         `json:"message,omitempty"`
	Name    string         `json:"name,omitempty"`
	Email   string         `json:"email,omitempty"`
	History *model.History `json:"history,omitempty"`
	Review  *model.Review  `json:"review,omitempty"`
}

func (s *Server) getCustomerhandler(c echo.Context) error {
	customerId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(404, &getCustomerResponse{Message: "id is invalid"})
	}
	ctx := context.Background()
	customer, err := s.customerStore.GetCustomer(ctx, customerId)
	if err != nil {
		return c.JSON(404, &getCustomerResponse{Message: "Not Found"})
	}
	historyErrCh := make(chan error, 1)
	reviewErrCh := make(chan error, 1)
	var h *model.History
	var r *model.Review

	go func() {
		var err error
		h, err = s.historiesClient.GetCustomerHistory(ctx, customerId)
		historyErrCh <- err
	}()

	go func() {
		var err error
		r, err = s.reviewsClient.GetCustomerReview(ctx, customerId)
		reviewErrCh <- err
	}()

	historyErr := <-historyErrCh
	if historyErr != nil {
		h = &model.History{Message: fmt.Sprint("failed to get history")}
	}
	reviewErr := <-reviewErrCh
	if reviewErr != nil {
		r = &model.Review{Message: fmt.Sprint("failed to get review")}
	}

	return c.JSON(200, &getCustomerResponse{
		Name:    customer.Name,
		Email:   customer.Email,
		History: h,
		Review:  r,
	})
}
