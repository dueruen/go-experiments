package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dueruen/go-experiments/go-kit/chat/pkg/storage/json"
	"github.com/dueruen/go-experiments/go-kit/chat/pkg/transport"
	pb "github.com/dueruen/go-experiments/go-kit/chat/pkg/transport/pb"
	"github.com/dueruen/go-experiments/go-kit/chat/pkg/writting"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	grpctransport "github.com/dueruen/go-experiments/go-kit/chat/pkg/transport/grpc"
)

type Type int

const (
	JSON Type = iota
	port      = ":50051"
)

func main() {
	storageType := JSON // this could be a flag; hardcoded here for simplicity

	var writter writting.Service
	//var lister listing.Service

	logger := makeLogger()

	switch storageType {
	case JSON:
		storage, _ := json.NewStorage()
		writter = writting.NewService(storage)
	}

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(writter)
	}

	var (
		ocTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{ocTracing}
		chatService     = grpctransport.NewGRPCServer(endpoints, serverOptions, logger)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	var g group.Group
	{
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", port)
			pb.RegisterChatServiceServer(grpcServer, chatService)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}

	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	level.Error(logger).Log("exit", g.Run())
}

func makeLogger() log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "chat",
			"ts", log.DefaultTimestampUTC,
			"clr", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	return logger
}
