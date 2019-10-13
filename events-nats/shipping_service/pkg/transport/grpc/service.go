package grpc

import (
	"context"

	pb "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/pb"
	"github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	reserveInventory kitgrpc.Handler
	shipOrder        kitgrpc.Handler
}

func NewGRPCServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.ShippingServiceServer {
	return &server{
		reserveInventory: kitgrpc.NewServer(
			endpoints.ReserveInventory, decodeReserveInventoryRequest, encodeReserveInventoryResponse,
		),
		shipOrder: kitgrpc.NewServer(
			endpoints.ShipOrder, decodeShipOrderRequest, encodeShipOrderResponse,
		),
	}
}

func (srv *server) ReserveInventory(ctx context.Context, req *pb.ReserveInventoryRequest) (*pb.ReserveInventoryResponse, error) {
	_, rep, err := srv.reserveInventory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ReserveInventoryResponse), nil
}

func decodeReserveInventoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.ReserveInventoryRequest), nil
}

func encodeReserveInventoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ReserveInventoryResponse), nil
}

func (srv *server) ShipOrder(ctx context.Context, req *pb.ShipOrderRequest) (*pb.ShipOrderResponse, error) {
	_, rep, err := srv.shipOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ShipOrderResponse), nil
}

func decodeShipOrderRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.ShipOrderRequest), nil
}

func encodeShipOrderResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ShipOrderResponse), nil
}
