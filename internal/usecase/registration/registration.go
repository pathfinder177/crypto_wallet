package registration

import (
	"context"
	"main/internal/entity"
	"main/internal/repo/persistent"
)

type UseCase struct {
	repo persistent.PersistentRepo
}

func New(r persistent.PersistentRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) Register(ctx context.Context, reg entity.Registration) (bool, error) {
	return uc.repo.Create(ctx, reg)
}
