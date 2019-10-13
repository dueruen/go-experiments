package transport

import (
	"context"

	pb "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/pb"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ReserveInventory endpoint.Endpoint
	ShipOrder        endpoint.Endpoint
}

func MakeEndpoints(srv pb.ShippingServiceServer) Endpoints {
	return Endpoints{
		ReserveInventory: makeReserveInventoryEndpoint(srv),
		ShipOrder:        makeShipOrderEndpoint(srv),
	}
}

func makeReserveInventoryEndpoint(srv pb.ShippingServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ReserveInventoryRequest)
		res, _ := srv.ReserveInventory(ctx, req)
		return res, nil
	}
}

func makeShipOrderEndpoint(srv pb.ShippingServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ShipOrderRequest)
		res, _ := srv.ShipOrder(ctx, req)
		return res, nil
	}
}
