package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type Repository struct {
	vessels []*pb.Vessel
}

// FindAvailable - checks a specificaiton against a map of vessels,
// return vaild vessel if capacity and weight are below vessel's.
func (repo *Repository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

// gRPC handler
type Service struct {
	repo repository
}

func (s *Service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the repsonse message type
	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "v-001", Name: "Noah Ork", MaxWeight: 200000, Capacity: 500},
	}
	repo := &Repository{vessels}

	srv := micro.NewService(
		micro.Name("micro.vessel.service"),
	)
	srv.Init()

	// Register handler
	pb.RegisterVesselServiceHandler(srv.Server(), &Service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
