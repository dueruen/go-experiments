package main

import (
	"fmt"
	"net"

	pb "github.com/dueruen/go-experiments/go-graphql/service/house/gen/proto"
	om "github.com/dueruen/go-experiments/go-graphql/service/house/pkg"
	"github.com/dueruen/go-experiments/go-graphql/service/house/pkg/storage/json"
	"github.com/dueruen/go-experiments/go-graphql/service/house/pkg/transport"
	grpctransport "github.com/dueruen/go-experiments/go-graphql/service/house/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

func main() {
	storage, _ := json.NewStorage()
	service := om.NewService(storage)

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		houseService    = grpctransport.NewGrpcServer(transport.MakeEndpoints(service), serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterHouseServiceServer(grpcServer, houseService)
	fmt.Printf("House service is listening on port %s...\n", port)

	err := grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
