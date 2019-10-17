package main

import (
	"fmt"
	"net"

	pb "github.com/dueruen/go-experiments/go-graphql/service/user/gen/proto"
	om "github.com/dueruen/go-experiments/go-graphql/service/user/pkg"
	"github.com/dueruen/go-experiments/go-graphql/service/user/pkg/storage/json"
	"github.com/dueruen/go-experiments/go-graphql/service/user/pkg/transport"
	grpctransport "github.com/dueruen/go-experiments/go-graphql/service/user/pkg/transport/grpc"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	storage, _ := json.NewStorage()
	service := om.NewService(storage)

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		userService     = grpctransport.NewGrpcServer(transport.MakeEndpoints(service), serverOptions)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	pb.RegisterUserServiceServer(grpcServer, userService)
	fmt.Printf("User service is listening on port %s...\n", port)

	err := grpcServer.Serve(grpcListener)
	fmt.Println("Serve() failed", err)
}
