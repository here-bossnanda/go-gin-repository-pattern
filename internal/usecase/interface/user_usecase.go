package usecase_interface

import (
	"context"

	"github.com/sekolahmu/boilerplate-go/internal/domain/entity"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, offset, limit int) ([]*entity.User, error)
}
