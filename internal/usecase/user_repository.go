package usecase

import (
	"context"
	"social-media-app/internal/entity"
)

type UserRepository interface {
	GetByName(ctx context.Context, name string) (entity.User, error)
	Save(ctx context.Context, u entity.User) (entity.User, error)
}
