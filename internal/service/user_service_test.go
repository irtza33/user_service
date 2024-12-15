package service

import (
    "context"
    "testing"

    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
    pb "user_service/proto/user"
)

type mockUserStore struct {
    mock.Mock
}

func (m *mockUserStore) GetUser(ctx context.Context, id int32) (*User, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*User), args.Error(1)
}

func TestGetUser(t *testing.T) {
    store := new(mockUserStore)
    service := NewUserService(store)

    store.On("GetUser", mock.Anything, int32(1)).Return(&User{ID: 1, Name: "Test"}, nil)

    resp, err := service.GetUser(context.Background(), &pb.GetUserRequest{UserId: 1})

    assert.NoError(t, err)
    assert.Equal(t, "Test", resp.Name)
    store.AssertExpectations(t)
}