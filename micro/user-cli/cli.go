package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	pb "github.com/shaqsnake/micro/user-service/proto/user"
)

func main() {
	srv := micro.NewService(
		micro.Name("micro.user.cli"),
		micro.Version("latest"),
	)

	srv.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("micro.user.service", microclient.DefaultClient)

	user := &pb.User{
		Name:     "shaqsnake",
		Email:    "shaqsnake@gmail.com",
		Password: "inputpasswd",
		Company:  "Google",
	}
	log.Printf("Add user: %v", user)

	// Call our user service
	r, err := client.Create(context.Background(), user)
	if err != nil {
		log.Fatalf("Could not create %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", user.Email, err)
	}
	log.Printf("Your access token is: %s\n", authResponse.Token)

	// Run the server
	// if err := srv.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	os.Exit(0)
}
