package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"social-media-app/internal/entity"
	"social-media-app/internal/usecase"
)

type UserRepository struct {
	db *sqlx.DB
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (entity.User, error) {
	q := "SELECT id, name FROM user WHERE name = :name"

	stmt, err := r.db.PrepareNamedContext(ctx, q)
	if err != nil {
		return entity.User{}, fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	var u entity.User
	err = stmt.GetContext(ctx, &u, map[string]any{"name": name})
	if err != nil {
		return entity.User{}, fmt.Errorf("get context: %w", err)
	}

	return u, nil
}

func (r *UserRepository) Save(ctx context.Context, u entity.User) (entity.User, error) {
	q := "INSERT INTO user (name) VALUES (:name) RETURNING id, name"

	stmt, err := r.db.PrepareNamedContext(ctx, q)
	if err != nil {
		return entity.User{}, fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, &u, u)
	if err != nil {
		return entity.User{}, fmt.Errorf("get context: %w", err)
	}

	return u, nil
}

var _ usecase.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}
