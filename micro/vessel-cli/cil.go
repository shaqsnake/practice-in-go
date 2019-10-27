package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

func main() {
	// Set-up vessel-cli service
	srv := micro.NewService(micro.Name("micro.vessel.cli"))
	srv.Init()

	client := pb.NewVesselServiceClient("micro.vessel.service", srv.Client())

	vessel := &pb.Vessel {
		Id: "v001",
		Name: "Noah Ork",
		Capacity: 500,
		MaxWeight: 200000,
	};

	// Call vessel service to create a new vessel
	res, err := client.Create(context.Background(), vessel)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Vessel created: %t, new vessel named: %s", res.Created, res.Vessel.Name)
}
