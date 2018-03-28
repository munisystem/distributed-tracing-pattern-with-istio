package reviews

import (
	"context"

	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/client"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/customers/model"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/api"
	"google.golang.org/grpc"
)

type clientImpl struct {
	cli pb.ReviewClient
}

func New(url string) (client.Reviews, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	cli := pb.NewReviewClient(conn)
	return &clientImpl{
		cli: cli,
	}, nil
}

func (c *clientImpl) GetCustomerReview(ctx context.Context, customerId int64) (*model.Review, error) {
	req := &pb.GetCustomerReviewRequest{CustomerId: customerId}
	res, err := c.cli.GetCustomerReview(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.Review{
		ItemName:    res.ItemName,
		Description: res.Description,
		Rating:      res.Rating,
	}, nil
}
