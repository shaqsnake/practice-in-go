package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("micro.vessel.service"),
	)
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri);
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("micro").Collection("vessel")
	repo := &VesselRepository {
		vesselCollection,
	}

	// Register handler
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
