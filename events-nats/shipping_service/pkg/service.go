package pkg

import (
	"context"
	"fmt"
	"log"
	"net"

	pub "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/events/pub"
	pb "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/pb"
	"github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/transport"
	grpctransport "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

const port = ":50051"

func Run() {
	srv := NewService()
	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(srv)
	}

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		shippingService = grpctransport.NewGRPCServer(endpoints, serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterShippingServiceServer(grpcServer, shippingService)

	fmt.Printf("Shipping service is listening on port %s...\n", port)

	eventSender, errSub := pub.NewEventSender("localhost:4222")
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Connection to NATS service made\n")

	srv.EventSender = eventSender

	err := grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}

type ShippingEventSender interface {
	SendEventInventoryReserved()
	SendEventOrderShipped()
}

type Service struct {
	EventSender ShippingEventSender
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) ReserveInventory(ctx context.Context, req *pb.ReserveInventoryRequest) (*pb.ReserveInventoryResponse, error) {
	res := new(pb.ReserveInventoryResponse)
	srv.EventSender.SendEventInventoryReserved()
	fmt.Printf("Called ReserveInventory\n")
	return res, nil
}

func (srv *Service) ShipOrder(ctx context.Context, req *pb.ShipOrderRequest) (*pb.ShipOrderResponse, error) {
	res := new(pb.ShipOrderResponse)
	srv.EventSender.SendEventOrderShipped()
	fmt.Printf("Called ShipOrder\n")
	return res, nil
}
