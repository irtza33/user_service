package repository

import (
    "context"
    "user_service/internal/domain"
)

type UserStore interface {
    GetUser(ctx context.Context, id int32) (*domain.User, error)
    CreateUser(ctx context.Context, name string) (int32, error)
    DeleteUser(ctx context.Context, id int32) error
}