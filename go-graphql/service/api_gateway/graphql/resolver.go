package graphql

import (
	"context"

	pb "github.com/dueruen/go-experiments/go-graphql/service/api_gateway/gen/proto"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserClient  pb.UserServiceClient
	HouseClient pb.HouseServiceClient
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *NewUser) (*User, error) {
	res, err := r.UserClient.CreateUser(ctx, &pb.CreateUserRequest{
		FirstName: input.Firstname,
		LastName:  input.Lastname,
		Age:       int32(input.Age),
	})
	if err != nil {
		return nil, err
	}
	return &User{
		Firstname: res.User.FirstName,
		Lastname:  res.User.LastName,
		Age:       int(res.User.Age),
		ID:        res.User.ID,
	}, nil
}
func (r *mutationResolver) CreateHouse(ctx context.Context, input *NewHouse) (*House, error) {
	res, err := r.HouseClient.CreateHouse(ctx, &pb.CreateHouseRequest{
		Address: input.Address,
		OwnerID: input.OwnerID,
		Age:     int32(input.Age),
	})
	if err != nil {
		return nil, err
	}
	return &House{
		Address: res.House.Address,
		OwnerID: res.House.OwnerID,
		Age:     int(res.House.Age),
		ID:      res.House.ID,
	}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	res, err := r.UserClient.ListAllUsers(ctx, &pb.ListAllUsersRequest{})
	if err != nil {
		return nil, err
	}
	list := []*User{}
	for _, u := range res.List {
		list = append(list, &User{
			Firstname: u.FirstName,
			Lastname:  u.LastName,
			Age:       int(u.Age),
			ID:        u.ID,
		})
	}
	return list, nil
}

func (r *queryResolver) Houses(ctx context.Context) ([]*House, error) {
	res, err := r.HouseClient.ListAllHouses(ctx, &pb.ListAllHousesRequest{})
	if err != nil {
		return nil, err
	}
	list := []*House{}
	for _, h := range res.List {
		list = append(list, &House{
			Address: h.Address,
			OwnerID: h.OwnerID,
			Age:     int(h.Age),
			ID:      h.ID,
		})
	}
	return list, nil
}
