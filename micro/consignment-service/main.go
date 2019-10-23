package main

import (
	"sync"
	"context"
	"net"
	"log"

	// Import the generated protobuf code
	pb "micro/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository - Dummy repository, simulates the use of datastore of some kind.
type Repository struct {
	mu          sync.RWMutex
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

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition.
type service struct {
	repo repository
}

// CreateConsignment
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// Save our consignment
	consignments, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching the 'Response' message we created in our protobuf definition.
	return &pb.Response{Created: true, Consignment: consignments}, nil
}

func main() {
	repo := &Repository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server.
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
