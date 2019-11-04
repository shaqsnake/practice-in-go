package main

import (
	"log"

	"github.com/micro/go-micro"
	pb "github.com/shaqsnake/micro/user-service/proto/user"
)

const topic = "user.created"

func main() {
	// Connect to database
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Migrate the user stuct int database.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Set-up user service
	srv := micro.NewService(
		micro.Name("micro.user.service"),
		micro.Version("latest"),
	)
	srv.Init()

	// Get publisher
	publisher := micro.NewPublisher(topic, srv.Client())

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &handler{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
