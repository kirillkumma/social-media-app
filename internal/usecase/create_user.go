package usecase

import (
	"context"
	"social-media-app/internal/entity"
)

type CreateUserInputDTO struct {
	Name string `json:"name"`
}

type CreateUserInteractor interface {
	Execute(ctx context.Context, p CreateUserInputDTO) (entity.User, error)
}

type CreateUserUseCase struct {
	userRepo UserRepository
}

func (u *CreateUserUseCase) Execute(ctx context.Context, p CreateUserInputDTO) (entity.User, error) {
	_, err := u.userRepo.GetByName(ctx, p.Name)
	switch err {
	case ErrNotFound:
		break
	case nil:
		return entity.User{}, entity.NewError("this name is already used", entity.ErrCodeBadInput)
	default:
		return entity.User{}, entity.NewError(err.Error(), entity.ErrCodeInternal)
	}

	user, err := entity.NewUser(p.Name)
	if err != nil {
		return entity.User{}, entity.NewError(err.Error(), entity.ErrCodeBadInput)
	}

	user, err = u.userRepo.Save(ctx, user)
	if err != nil {
		return entity.User{}, entity.NewError(err.Error(), entity.ErrCodeInternal)
	}

	return user, nil
}

var _ CreateUserInteractor = (*CreateUserUseCase)(nil)

func NewCreateUserUseCase(userRepo UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo}
}
