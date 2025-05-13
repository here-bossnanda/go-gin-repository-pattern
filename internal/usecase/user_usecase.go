package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sekolahmu/boilerplate-go/internal/domain/entity"
	repoInterface "github.com/sekolahmu/boilerplate-go/internal/repository/interface"
	"github.com/sekolahmu/boilerplate-go/internal/usecase/usecase_interface"
)

type userUseCase struct {
	userRepo repoInterface.Repository[entity.User]
}

// NewUserUseCase membuat instance baru dari UserUseCase
func NewUserUseCase(userRepo repoInterface.Repository[entity.User]) usecase_interface.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return uc.userRepo.Create(ctx, user)
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *userUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	existingUser, err := uc.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return entity.ErrUserNotFound
	}

	user.UpdatedAt = time.Now()
	return uc.userRepo.Update(ctx, user)
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id string) error {
	existingUser, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return entity.ErrUserNotFound
	}

	return uc.userRepo.Delete(ctx, id)
}

func (uc *userUseCase) ListUsers(ctx context.Context, offset, limit int) ([]*entity.User, error) {
	return uc.userRepo.List(ctx, offset, limit)
}
