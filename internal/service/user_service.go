package service

import (
    "context"
    "errors"
    "log"

    "github.com/irtza33/user_service/internal/repository"
    "github.com/irtza33/user_service/pkg/logger"
    "github.com/irtza33/user_service/proto/user"
)

type UserService struct {
    userStore repository.UserStore
    logger     logger.Logger
}

func NewUserService(userStore repository.UserStore, logger logger.Logger) *UserService {
    return &UserService{
        userStore: userStore,
        logger:     logger,
    }
}

func (s *UserService) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
    if req.UserId <= 0 {
        s.logger.Error("Invalid user ID")
        return nil, errors.New("invalid user ID")
    }

    name, err := s.userStore.GetUser(req.UserId)
    if err != nil {
        s.logger.Error("Error fetching user: %v", err)
        return nil, err
    }

    return &user.GetUserResponse{Name: name}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
    if req.Name == "" {
        s.logger.Error("Invalid name")
        return nil, errors.New("invalid name")
    }

    userId, err := s.userStore.CreateUser(req.Name)
    if err != nil {
        s.logger.Error("Error creating user: %v", err)
        return nil, err
    }

    return &user.CreateUserResponse{UserId: userId}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
    if req.UserId <= 0 {
        s.logger.Error("Invalid user ID")
        return nil, errors.New("invalid user ID")
    }

    err := s.userStore.DeleteUser(req.UserId)
    if err != nil {
        s.logger.Error("Error deleting user: %v", err)
        return nil, err
    }

    return &user.DeleteUserResponse{Confirmation: "User deleted successfully"}, nil
}