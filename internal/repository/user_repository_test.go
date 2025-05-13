package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sekolahmu/boilerplate-go/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	// Gunakan database test
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/boilerplate_test?sslmode=disable")
	require.NoError(t, err)

	// Bersihkan tabel sebelum test
	_, err = db.Exec("TRUNCATE TABLE users")
	require.NoError(t, err)

	return db
}

func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()

	user := &entity.User{
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Create(ctx, user)
	require.NoError(t, err)
	assert.NotEmpty(t, user.ID)

	// Verifikasi user tersimpan di database
	var savedUser entity.User
	err = db.QueryRow("SELECT id, email, name, password, created_at, updated_at FROM users WHERE id = $1", user.ID).
		Scan(&savedUser.ID, &savedUser.Email, &savedUser.Name, &savedUser.Password, &savedUser.CreatedAt, &savedUser.UpdatedAt)
	require.NoError(t, err)
	assert.Equal(t, user.Email, savedUser.Email)
	assert.Equal(t, user.Name, savedUser.Name)
}

func TestUserRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()

	// Buat user untuk test
	user := &entity.User{
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := repo.Create(ctx, user)
	require.NoError(t, err)

	// Test GetByID
	foundUser, err := repo.GetByID(ctx, user.ID)
	require.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Name, foundUser.Name)

	// Test GetByID dengan ID yang tidak ada
	notFoundUser, err := repo.GetByID(ctx, "non-existent-id")
	require.NoError(t, err)
	assert.Nil(t, notFoundUser)
}

func TestUserRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()

	// Buat user untuk test
	user := &entity.User{
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := repo.Create(ctx, user)
	require.NoError(t, err)

	// Update user
	user.Name = "Updated Name"
	user.Email = "updated@example.com"
	user.UpdatedAt = time.Now()
	err = repo.Update(ctx, user)
	require.NoError(t, err)

	// Verifikasi update
	updatedUser, err := repo.GetByID(ctx, user.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated Name", updatedUser.Name)
	assert.Equal(t, "updated@example.com", updatedUser.Email)
}

func TestUserRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()

	// Buat user untuk test
	user := &entity.User{
		Email:     "test@example.com",
		Name:      "Test User",
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := repo.Create(ctx, user)
	require.NoError(t, err)

	// Delete user
	err = repo.Delete(ctx, user.ID)
	require.NoError(t, err)

	// Verifikasi user sudah dihapus
	deletedUser, err := repo.GetByID(ctx, user.ID)
	require.NoError(t, err)
	assert.Nil(t, deletedUser)
}

func TestUserRepository_List(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)
	ctx := context.Background()

	// Buat beberapa user untuk test
	users := []*entity.User{
		{
			Email:     "user1@example.com",
			Name:      "User 1",
			Password:  "password123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "user2@example.com",
			Name:      "User 2",
			Password:  "password123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Email:     "user3@example.com",
			Name:      "User 3",
			Password:  "password123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		err := repo.Create(ctx, user)
		require.NoError(t, err)
	}

	// Test List dengan pagination
	list, err := repo.List(ctx, 0, 2)
	require.NoError(t, err)
	assert.Len(t, list, 2)

	// Test List dengan offset
	list, err = repo.List(ctx, 2, 2)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}
