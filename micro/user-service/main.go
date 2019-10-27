package main

import (
	"log"

	pb "github.com/shaqsnake/micro/user-service/proto/user"
	"github.com/micro/go-micro"
)

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
	
	// Set-up user service
	srv := micro.NewService(
		micro.Name("micro.user.service"),
	)
	srv.Init()

	// Register handler
	pb.RegisterUserServiceHandler(srv.Server(), &handler{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}