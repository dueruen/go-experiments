package transport

import (
	"context"

	writting "github.com/dueruen/go-experiments/go-kit/chat/pkg/writting"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	WriteToChat endpoint.Endpoint
	//GetChat     endpoint.Endpoint
}

func MakeEndpoints(service writting.Service) Endpoints {
	return Endpoints{
		WriteToChat: makeWriteToChatEndpoint(service),
	}
}

func makeWriteToChatEndpoint(service writting.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WriteToChatRequest)
		res, _ := service.WriteToChat(ctx, req.ChatItem)
		return WriteToChatResponse{ChatItem: res}, nil
	}
}
