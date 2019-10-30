package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	pb "github.com/shaqsnake/micro/consignment-service/proto/consignment"
	userPb "github.com/shaqsnake/micro/user-service/proto/user"
	vesselPb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	// Set-up a new service.
	srv := micro.NewService(
		micro.Name("micro.consignment.service"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)

	// Init will parse the command line flags.
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("micro").Collection("consignment")

	repo := &MongoRepository{consignmentCollection}
	vesselClient := vesselPb.NewVesselServiceClient("micro.vessel.service", srv.Client())

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &handler{repo, vesselClient})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("No auth meta_data found in request")
		}

		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth
		authClient := userPb.NewUserServiceClient("micro.user.service", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userPb.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, res)
		return err
	}
}
