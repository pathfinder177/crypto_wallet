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
	if _, err := uc.repo.Create(ctx, reg); err != nil {
		return false, err
	}

	return true, nil
}
