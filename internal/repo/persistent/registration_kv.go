package persistent

import (
	"context"
	"main/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type RegistrationRepo struct {
	repo map[string]string
}

func NewRegistrationRepo(m map[string]string) *RegistrationRepo {
	return &RegistrationRepo{m}
}

func (repo *RegistrationRepo) Create(ctx context.Context, reg entity.Registration) (bool, error) {
	uname := reg.Username
	upass := reg.Password

	if _, ok := repo.repo[uname]; ok {
		return false, errRegUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(upass), bcrypt.DefaultCost)
	if err != nil {
		return false, errRegBcryptGenFromPassword
	}

	repo.repo[uname] = string(passwordHash)

	return true, nil
}
