package transport

import (
	"context"

	om "github.com/dueruen/go-experiments/go-graphql/service/house/pkg"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateHouse   endpoint.Endpoint
	ListAllHouses endpoint.Endpoint
}

func MakeEndpoints(svc om.HouseService) Endpoints {
	return Endpoints{
		CreateHouse:   makeCreateHouseEndpoint(svc),
		ListAllHouses: makeListAllHousesEndpoint(svc),
	}
}

func makeCreateHouseEndpoint(svc om.HouseService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(om.CreateHouseRequest)
		res, _ := svc.CreateHouse(req)
		return res, nil
	}
}

func makeListAllHousesEndpoint(svc om.HouseService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(om.ListAllHousesRequest)
		res, _ := svc.ListAllHouses(req)
		return res, nil
	}
}
