package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/user-service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	repo         Repository
	tokenService Authable
	publisher    micro.Publisher
}

func (h *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// Gen hashed password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := h.repo.Create(req); err != nil {
		return err
	}
	res.User = req

	if err := h.publisher.Publish(ctx, req); err != nil {
		return err
	}
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
	log.Println("Logging in with:", req.Email, req.Password)
	user, err := h.repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := h.tokenService.Encode(user)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	// Decode token
	claims, err := h.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("Invalid user")
	}

	res.Valid = true
	return nil
}
