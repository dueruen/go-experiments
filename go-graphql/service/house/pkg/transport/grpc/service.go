package grpc

import (
	"context"

	pb "github.com/dueruen/go-experiments/go-graphql/service/house/gen/proto"
	om "github.com/dueruen/go-experiments/go-graphql/service/house/pkg"
	"github.com/dueruen/go-experiments/go-graphql/service/house/pkg/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	createHouse   kitgrpc.Handler
	listAllHouses kitgrpc.Handler
}

func NewGrpcServer(endpoins transport.Endpoints, options []kitgrpc.ServerOption) pb.HouseServiceServer {
	return &server{
		createHouse:   kitgrpc.NewServer(endpoins.CreateHouse, decodeCreateHouseRequest, encodeCreateHouseResponse),
		listAllHouses: kitgrpc.NewServer(endpoins.ListAllHouses, decodeListAllHousesRequest, encodeListAllHousesResponse),
	}
}

func (server *server) CreateHouse(ctx context.Context, req *pb.CreateHouseRequest) (*pb.CreateHouseResponse, error) {
	_, rep, err := server.createHouse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateHouseResponse), nil
}

func decodeCreateHouseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateHouseRequest)
	return om.CreateHouseRequest{
		Address: req.Address,
		OwnerID: req.OwnerID,
		Age:     req.Age,
	}, nil
}

func encodeCreateHouseResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(om.CreateHouseResponse)
	return &pb.CreateHouseResponse{
		House: &pb.House{
			Address: res.House.Address,
			OwnerID: res.House.OwnerID,
			Age:     res.House.Age,
			ID:      res.House.ID,
		},
	}, nil
}

func (server *server) ListAllHouses(ctx context.Context, req *pb.ListAllHousesRequest) (*pb.ListAllHousesResponse, error) {
	_, rep, err := server.listAllHouses.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListAllHousesResponse), nil
}

func decodeListAllHousesRequest(_ context.Context, request interface{}) (interface{}, error) {
	return om.ListAllHousesRequest{}, nil
}

func encodeListAllHousesResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(om.ListAllHousesResponse)
	list := []*pb.House{}
	for _, h := range res.Houses {
		list = append(list, &pb.House{
			Address: h.Address,
			OwnerID: h.OwnerID,
			Age:     h.Age,
			ID:      h.ID,
		})
	}
	return &pb.ListAllHousesResponse{
		List: list,
	}, nil
}
