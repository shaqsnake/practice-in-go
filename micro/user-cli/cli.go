package main

import (
	"log"
	"os"
	"context"
	
	pb "github.com/shaqsnake/micro/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("micro.user.service", microclient.DefaultClient)

	// Default our flags
	srv := micro.NewService(
		micro.Flags(
			cli.StringFlag {
				Name: "name",
				Usage: "Your full name",
			},
			cli.StringFlag {
				Name: "email",
				Usage: "Your email",
			},
			cli.StringFlag {
				Name: "password",
				Usage: "Your password",
			},
			cli.StringFlag {
				Name: "company",
				Usage: "Your company",
			},
		),
	)

	// Start as service
	srv.Init(
		micro.Action(func(c *cli.Context) {
			log.Printf("%v", c.Args())
			// name := c.String("name")
			// email := c.String("email")
			// password := c.String("password")
			// company := c.String("company")

			user := &pb.User{
				Name: "shaqsnake",
				Email: "shaqsnake@gmail.com",
				Password: "inputpasswd",
				Company: "Google",
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

			os.Exit(0)
		}),
	)

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}