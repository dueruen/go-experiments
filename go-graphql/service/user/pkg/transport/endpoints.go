package transport

import (
	"context"

	om "github.com/dueruen/go-experiments/go-graphql/service/user/pkg"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser   endpoint.Endpoint
	ListAllUsers endpoint.Endpoint
}

func MakeEndpoints(svc om.UserService) Endpoints {
	return Endpoints{
		CreateUser:   makeCreateUserEndpoint(svc),
		ListAllUsers: makeListAllUsersEndpoint(svc),
	}
}

func makeCreateUserEndpoint(svc om.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(om.CreateUserRequest)
		res, _ := svc.CreateUser(req)
		return res, nil
	}
}

func makeListAllUsersEndpoint(svc om.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(om.ListAllUsersRequest)
		res, _ := svc.ListAllUsers(req)
		return res, nil
	}
}
