package persistent

import (
	"context"
	"errors"
	"main/internal/entity"
)

type LoginRepo struct {
	repo map[string]string
}

func NewLoginRepo(m map[string]string) *LoginRepo {
	return &LoginRepo{m}
}

func (repo *LoginRepo) Read(ctx context.Context, reg entity.Registration) (bool, error) {
	uname := reg.Username
	if _, ok := repo.repo[uname]; !ok {
		return false, errors.New("no such user, please register")
	}

	return true, nil
}
