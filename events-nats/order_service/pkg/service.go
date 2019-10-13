package pkg

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dueruen/go-experiments/events-nats/order_service/pkg/events/sub"
	pb "github.com/dueruen/go-experiments/events-nats/order_service/pkg/pb"
	"github.com/dueruen/go-experiments/events-nats/order_service/pkg/transport"
	grpctransport "github.com/dueruen/go-experiments/events-nats/order_service/pkg/transport/grpc"
	shippingPb "github.com/dueruen/go-experiments/events-nats/shipping_service/pkg/pb"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

const port = ":50052"

func Run() {
	srv := NewService()
	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(srv)
	}

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		orderService    = grpctransport.NewGRPCServer(endpoints, serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterOrderServiceServer(grpcServer, orderService)

	fmt.Printf("Order service is listening on port %s...\n", port)

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to shipping service %v", err)
	}
	defer cc.Close()

	c := shippingPb.NewShippingServiceClient(cc)
	fmt.Printf("Connection to shipping service made\n")

	srv.ShippingClient = c

	errSub := sub.Listen("localhost:4222", c)
	if errSub != nil {
		log.Fatalf("Could not connect to NATS %v", errSub)
	}
	fmt.Printf("Connection to NATS service made\n")

	errGRPC := grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", errGRPC)
}

type Service struct {
	ShippingClient shippingPb.ShippingServiceClient
}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	res := new(pb.CreateOrderResponse)
	srv.ShippingClient.ReserveInventory(context.Background(), &shippingPb.ReserveInventoryRequest{})
	return res, nil
}

func (srv *Service) GetShippedOrders(ctx context.Context, req *pb.GetShippedOrdersRequest) (*pb.GetShippedOrdersResponse, error) {
	res := new(pb.GetShippedOrdersResponse)
	res.Amount = sub.GetOrdersShipped()
	return res, nil
}
