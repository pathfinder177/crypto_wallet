package persistent

import (
	"context"
	"errors"
	"main/internal/entity"
)

type RegistrationRepo struct {
	repo map[string]string
}

func NewRegistrationRepo(m map[string]string) *RegistrationRepo {
	return &RegistrationRepo{m}
}

func (repo *RegistrationRepo) Create(ctx context.Context, reg entity.Registration) (bool, error) {
	uname := reg.Username
	upass := reg.HashedPassword

	if _, ok := repo.repo[uname]; ok {
		return false, errors.New("user exists")
	}

	repo.repo[uname] = upass
	return true, nil
}
