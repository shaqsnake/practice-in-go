package main

import (
	"context"

	pb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

// gRPC handler
type handler struct {
	repo Repository
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := h.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the repsonse message type
	res.Vessel = vessel
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := h.repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}