package handler

import (
	"context"

	"github.com/sikemausa/micro-service-example/internal/domain"
	"github.com/sikemausa/micro-service-example/internal/service"
	v1 "github.com/sikemausa/micro-service-example/pb/v1"
)

type UserServiceServer struct {
	service *service.UserService
	v1.UnimplementedUserServiceServer
}

func NewUserServiceServer(service *service.UserService) *UserServiceServer {
	return &UserServiceServer{service: service}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	user := domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	createdUser, err := s.service.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &v1.CreateUserResponse{
		User: &v1.User{
			Id:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
	}
	return response, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	user, err := s.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	response := &v1.GetUserResponse{
		User: &v1.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return response, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	user := domain.User{
		ID:    req.Id,
		Name:  req.User.Name,
		Email: req.User.Email,
	}

	user, err := s.service.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	response := &v1.UpdateUserResponse{
		User: &v1.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return response, nil
}

func (s *UserServiceServer) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	err := s.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	response := &v1.DeleteUserResponse{
		Message: "User deleted successfully",
	}
	return response, nil
}

func (s *UserServiceServer) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	users, err := s.service.List(ctx)
	if err != nil {
		return nil, err
	}

	var userList []*v1.User
	for _, user := range users {
		userList = append(userList, &v1.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	response := &v1.ListUsersResponse{
		Users: userList,
	}
	return response, nil
}
