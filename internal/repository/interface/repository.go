package repository

import (
	"context"
)

// Repository adalah interface dasar untuk semua repository
type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, offset, limit int) ([]*T, error)
}

// Transactional adalah interface untuk mendukung transaksi database
type Transactional interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
