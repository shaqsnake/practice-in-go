package main

import (
	"context"
	"log"

	pb "github.com/shaqsnake/micro/consignment-service/proto/consignment"
	vesselPb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

type handler struct {
	repo Repository
	vesselClient vesselPb.VesselServiceClient
}

// CreateConsignment -
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Call a vessel service with consignment weight and capacity value
	vesselRes, err := s.vesselClient.FindAvailable(context.Background(), &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s \n", vesselRes.Vessel.Name)

	// Set the VesselId which from the vessel response
	req.VesselId = vesselRes.Vessel.Id

	// Save our consignment
	if err := s.repo.Create(req); err != nil {
		return err
	}

	// Return matching the 'Response' message we created in our protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments -
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
