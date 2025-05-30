package login

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

func (uc *UseCase) Login(ctx context.Context, login entity.Login) (bool, error) {
	return uc.repo.Read(ctx, login)
}
