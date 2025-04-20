package login

import (
	"context"
	"main/internal/entity"
	"main/internal/repo"
)

type UseCase struct {
	repo repo.LoginRepo
}

func New(r repo.LoginRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) Login(ctx context.Context, reg entity.Registration) (bool, error) {
	return uc.repo.Read(ctx, reg)
}
