package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/shaqsnake/micro/consignment-service/proto/consignment"
)

// Repository -
type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (mongoRepo *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := mongoRepo.collection.InsertOne(context.Background(), consignment)
	return err
}

// GetAll -
func (mongoRepo *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	cur, err := mongoRepo.collection.Find(context.Background(), nil, nil)
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}

	return consignments, err
}
