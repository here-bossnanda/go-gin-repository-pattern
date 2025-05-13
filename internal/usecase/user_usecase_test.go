package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/sekolahmu/boilerplate-go/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockUserRepository adalah mock untuk UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) List(ctx context.Context, offset, limit int) ([]*entity.User, error) {
	args := m.Called(ctx, offset, limit)
	return args.Get(0).([]*entity.User), args.Error(1)
}

func (m *MockUserRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)
	ctx := context.Background()

	user := &entity.User{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password123",
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*entity.User")).Return(nil)

	err := useCase.CreateUser(ctx, user)
	require.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotZero(t, user.CreatedAt)
	assert.NotZero(t, user.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)
	ctx := context.Background()

	expectedUser := &entity.User{
		ID:        "test-id",
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetByID", ctx, "test-id").Return(expectedUser, nil)

	user, err := useCase.GetUserByID(ctx, "test-id")
	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_UpdateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)
	ctx := context.Background()

	existingUser := &entity.User{
		ID:        "test-id",
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updatedUser := &entity.User{
		ID:        "test-id",
		Email:     "updated@example.com",
		Name:      "Updated User",
		Password:  "newpassword",
		CreatedAt: existingUser.CreatedAt,
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetByID", ctx, "test-id").Return(existingUser, nil)
	mockRepo.On("Update", ctx, mock.AnythingOfType("*entity.User")).Return(nil)

	err := useCase.UpdateUser(ctx, updatedUser)
	require.NoError(t, err)
	assert.NotEqual(t, existingUser.UpdatedAt, updatedUser.UpdatedAt)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)
	ctx := context.Background()

	existingUser := &entity.User{
		ID:        "test-id",
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetByID", ctx, "test-id").Return(existingUser, nil)
	mockRepo.On("Delete", ctx, "test-id").Return(nil)

	err := useCase.DeleteUser(ctx, "test-id")
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserUseCase_ListUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)
	ctx := context.Background()

	expectedUsers := []*entity.User{
		{
			ID:        "user1",
			Email:     "user1@example.com",
			Name:      "User 1",
			Password:  "password123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "user2",
			Email:     "user2@example.com",
			Name:      "User 2",
			Password:  "password123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("List", ctx, 0, 10).Return(expectedUsers, nil)

	users, err := useCase.ListUsers(ctx, 0, 10)
	require.NoError(t, err)
	assert.Equal(t, expectedUsers, users)

	mockRepo.AssertExpectations(t)
}
