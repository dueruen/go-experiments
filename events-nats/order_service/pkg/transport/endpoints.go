package transport

import (
	"context"

	pb "github.com/dueruen/go-experiments/events-nats/order_service/pkg/pb"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateOrder      endpoint.Endpoint
	GetShippedOrders endpoint.Endpoint
}

func MakeEndpoints(srv pb.OrderServiceServer) Endpoints {
	return Endpoints{
		CreateOrder:      makeCreateOrderEndpoint(srv),
		GetShippedOrders: makeGetShippedOrdersEndpoint(srv),
	}
}

func makeCreateOrderEndpoint(srv pb.OrderServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateOrderRequest)
		res, _ := srv.CreateOrder(ctx, req)
		return res, nil
	}
}

func makeGetShippedOrdersEndpoint(srv pb.OrderServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetShippedOrdersRequest)
		res, _ := srv.GetShippedOrders(ctx, req)
		return res, nil
	}
}
