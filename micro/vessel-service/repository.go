package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/shaqsnake/micro/vessel-service/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
}

// VesselRepository -
type VesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable - checks a specificaiton against vessel in db,
// return vaild vessel if capacity and weight are below vessel's.
func (vesselRepo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	filter := bson.D{{"capacity", bson.M{"$gte": spec.Capacity}}, {"maxweight", bson.M{"$gte": spec.MaxWeight}}}
	if err := vesselRepo.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		return nil, err
	}

	return vessel, nil
}

// Create a new vessel
func (vesselRepo *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := vesselRepo.collection.InsertOne(context.TODO(), vessel)
	return err
}