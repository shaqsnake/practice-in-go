package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/shaqsnake/micro/user-service/proto/user"
)

type Repository interface {
	Create(*pb.User) error
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetByEmailAndPassword(*pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (userRepo *UserRepository) Create(user *pb.User) error{
	if err := userRepo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := userRepo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := userRepo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := userRepo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}