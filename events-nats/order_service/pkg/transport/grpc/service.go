package grpc

import (
	"context"

	pb "github.com/dueruen/go-experiments/events-nats/order_service/pkg/pb"
	"github.com/dueruen/go-experiments/events-nats/order_service/pkg/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	createOrder      kitgrpc.Handler
	getShippedOrders kitgrpc.Handler
}

func NewGRPCServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.OrderServiceServer {
	return &server{
		createOrder: kitgrpc.NewServer(
			endpoints.CreateOrder, decodeCreateOrderRequest, encodeCreateOrderResponse,
		),
		getShippedOrders: kitgrpc.NewServer(
			endpoints.GetShippedOrders, decodeGetShippedOrdersRequest, encodeGetShippedOrdersResponse,
		),
	}
}

func (srv *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	_, rep, err := srv.createOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateOrderResponse), nil
}

func decodeCreateOrderRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.CreateOrderRequest), nil
}

func encodeCreateOrderResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.CreateOrderResponse), nil
}

func (srv *server) GetShippedOrders(ctx context.Context, req *pb.GetShippedOrdersRequest) (*pb.GetShippedOrdersResponse, error) {
	_, rep, err := srv.getShippedOrders.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetShippedOrdersResponse), nil
}

func decodeGetShippedOrdersRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.GetShippedOrdersRequest), nil
}

func encodeGetShippedOrdersResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.GetShippedOrdersResponse), nil
}
