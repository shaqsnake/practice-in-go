package main

import (
	"context"
	"log"
	"os"

	// Import the generated protobuf code
	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/consignment-service/proto/consignment"
	vesselPb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	// Set-up a new service.
	srv := micro.NewService(
		micro.Name("micro.consignment.service"),
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
