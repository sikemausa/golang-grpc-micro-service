package handler

import (
	"context"
	"fmt"

	"github.com/sikemausa/micro-service-example/internal/domain"
	"github.com/sikemausa/micro-service-example/internal/service"
	"github.com/sikemausa/micro-service-example/pkg/proto"
)

type UserServiceServer struct {
	service *service.UserService
	proto.UnimplementedUserServiceServer
}

func NewUserServiceServer(service *service.UserService) *UserServiceServer {
	return &UserServiceServer{service: service}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	user := domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	createdUser, err := s.service.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &proto.UserResponse{
		User: &proto.User{
			Id:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
	}
	return response, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	fmt.Println("GetUser")
	user, err := s.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	response := &proto.UserResponse{
		User: &proto.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return response, nil
}
