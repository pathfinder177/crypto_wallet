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
	upass := reg.HashedPassword

	if hash, ok := repo.repo[uname]; !ok {
		return false, errors.New("no such user, please register")
	} else if hash != upass {
		return false, errors.New("password does not match")
	}

	return true, nil
}
