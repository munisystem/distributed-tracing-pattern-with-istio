package server

import (
	"context"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/api"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/client"
	"github.com/munisystem/distributed-tracing-pattern-with-istio/src/reviews/store"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	itemClient  client.Items
	reviewStore store.ReviewStore
	logger      *zap.Logger
}

func New(itemClient client.Items, reviewStore store.ReviewStore, logger *zap.Logger) *Server {
	return &Server{
		itemClient:  itemClient,
		reviewStore: reviewStore,
		logger:      logger,
	}
}

func (s *Server) Run(port string) error {
	l, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpc_zap.ReplaceGrpcLogger(s.logger)
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(s.logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(s.logger),
		)),
	)
	pb.RegisterReviewServer(srv, s)
	return srv.Serve(l)
}

func (s *Server) GetCustomerReview(ctx context.Context, in *pb.GetCustomerReviewRequest) (*pb.GetCustomerReviewResponse, error) {
	r, err := s.reviewStore.GetCustomerReview(ctx, in.CustomerId)
	if err != nil {
		return nil, err
	}
	i, err := s.itemClient.GetItem(ctx, r.ItemId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCustomerReviewResponse{
		Id:          r.Id,
		CustomerId:  in.CustomerId,
		ItemName:    i.Name,
		Rating:      r.Rating,
		Description: r.Description,
	}, nil
}
