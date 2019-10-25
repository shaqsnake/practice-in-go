package main

import (
	"context"
	"log"
	"sync"

	// Import the generated protobuf code
	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/consignment-service/proto/consignment"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, simulates the use of datastore of some kind.
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition.
type Service struct {
	repo repository
}

// CreateConsignment
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the 'Response' message we created in our protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

// GetConsignments
func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	// Create a new service.
	srv := micro.NewService(
		micro.Name("micro.consignment.service"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &Service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
