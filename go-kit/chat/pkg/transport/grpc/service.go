package grpc

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/dueruen/go-experiments/go-kit/chat/pkg/transport"
	pb "github.com/dueruen/go-experiments/go-kit/chat/pkg/transport/pb"
	writting "github.com/dueruen/go-experiments/go-kit/chat/pkg/writting"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	writeToChat kitgrpc.Handler
	logger      log.Logger
}

func NewGRPCServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption, logger log.Logger) pb.ChatServiceServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &server{
		writeToChat: kitgrpc.NewServer(
			endpoints.WriteToChat, decodeWriteToChatRequest, encodeWriteToChatResponse, options...,
		),
		logger: logger,
	}
}

func (server *server) WriteToChat(ctx context.Context, req *pb.WriteToChatRequest) (*pb.WriteToChatResponse, error) {
	_, rep, err := server.writeToChat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.WriteToChatResponse), nil
}

func decodeWriteToChatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.WriteToChatRequest)
	return transport.WriteToChatRequest{
		ChatItem: writting.ChatItem{
			Auther:  req.Auther,
			Message: req.Message,
		},
	}, nil
}

func encodeWriteToChatResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.WriteToChatResponse)
	return &pb.WriteToChatResponse{
		Auther:  res.ChatItem.Auther,
		Message: res.ChatItem.Message,
		Id:      res.ChatItem.ID,
	}, nil
}
