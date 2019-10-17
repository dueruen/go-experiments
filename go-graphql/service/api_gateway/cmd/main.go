package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/99designs/gqlgen/handler"
	pb "github.com/dueruen/go-experiments/go-graphql/service/api_gateway/gen/proto"
	"github.com/dueruen/go-experiments/go-graphql/service/api_gateway/graphql"
	"google.golang.org/grpc"
)

const port = 4000

func main() {
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to user service %v", err)
	}
	defer userConn.Close()
	userService := pb.NewUserServiceClient(userConn)
	fmt.Printf("Connection to user service made\n")

	houseConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to house service %v", err)
	}
	defer houseConn.Close()
	houseService := pb.NewHouseServiceClient(houseConn)
	fmt.Printf("Connection to house service made\n")

	resolver := graphql.Resolver{
		UserClient:  userService,
		HouseClient: houseService,
	}

	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		graphql.NewExecutableSchema(graphql.Config{Resolvers: &resolver}),
		handler.RecoverFunc(func(ctx context.Context, err interface{}) error {
			// send this panic somewhere
			log.Print(err)
			debug.PrintStack()
			return errors.New("user message on panic")
		}),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
