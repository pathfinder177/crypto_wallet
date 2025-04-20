package registration

import (
	"context"
	"main/internal/entity"
	"main/internal/repo"
)

type UseCase struct {
	repo repo.RegistrationRepo
}

func New(r repo.RegistrationRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) Register(ctx context.Context, reg entity.Registration) (bool, error) {
	return uc.repo.Create(ctx, reg)
}
