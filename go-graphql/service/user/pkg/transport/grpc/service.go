package grpc

import (
	"context"

	pb "github.com/dueruen/go-experiments/go-graphql/service/user/gen/proto"
	om "github.com/dueruen/go-experiments/go-graphql/service/user/pkg"
	"github.com/dueruen/go-experiments/go-graphql/service/user/pkg/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	createUser   kitgrpc.Handler
	listAllUsers kitgrpc.Handler
}

func NewGrpcServer(endpoins transport.Endpoints, options []kitgrpc.ServerOption) pb.UserServiceServer {
	return &server{
		createUser:   kitgrpc.NewServer(endpoins.CreateUser, decodeCreateUserRequest, encodeCreateUserResponse),
		listAllUsers: kitgrpc.NewServer(endpoins.ListAllUsers, decodeListAllUsersRequest, encodeListAllUsersResponse),
	}
}

func (server *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, rep, err := server.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserResponse), nil
}

func decodeCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserRequest)
	return om.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
	}, nil
}

func encodeCreateUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(om.CreateUserResponse)
	return &pb.CreateUserResponse{
		User: &pb.User{
			FirstName: res.User.FirstName,
			LastName:  res.User.LastName,
			Age:       res.User.Age,
			ID:        res.User.ID,
		},
	}, nil
}

func (server *server) ListAllUsers(ctx context.Context, req *pb.ListAllUsersRequest) (*pb.ListAllUsersResponse, error) {
	_, rep, err := server.listAllUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListAllUsersResponse), nil
}

func decodeListAllUsersRequest(_ context.Context, request interface{}) (interface{}, error) {
	return om.ListAllUsersRequest{}, nil
}

func encodeListAllUsersResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(om.ListAllUsersResponse)
	list := []*pb.User{}
	for _, u := range res.Users {
		list = append(list, &pb.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			ID:        u.ID,
		})
	}
	return &pb.ListAllUsersResponse{
		List: list,
	}, nil
}
