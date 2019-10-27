package main

import (
	"context"
	
	pb "github.com/shaqsnake/micro/user-service/proto/user"
)

type handler struct {
	repo Repository
	// tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := h.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := h.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := h.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "TestAuth"
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
