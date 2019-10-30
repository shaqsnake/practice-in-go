package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/metadata"
	pb "github.com/shaqsnake/micro/consignment-service/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecting file and token."))
	}
	file := os.Args[1]
	token := os.Args[2]

	client := pb.NewShippingServiceClient("micro.consignment.service", microclient.DefaultClient)

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create a new context which contains the token.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
