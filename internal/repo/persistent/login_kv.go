package persistent

import (
	"context"
	"main/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type LoginRepo struct {
	repo map[string]string
}

func NewLoginRepo(m map[string]string) *LoginRepo {
	return &LoginRepo{m}
}

func (repo *LoginRepo) Read(ctx context.Context, login entity.Login) (bool, error) {
	uname := login.Username
	upass := login.Password

	hash, prs := repo.repo[uname]

	if !prs {
		return false, errLoginNoUser
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(upass)); err != nil {
		return false, errLoginNoMatchPassword
	}

	return true, nil
}
