package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/user-service/proto/user"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("micro.email.service"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
