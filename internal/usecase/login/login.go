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
	if _, err := uc.repo.Read(ctx, reg); err != nil {
		return false, err
	}

	return true, nil
}
