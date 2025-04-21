package persistent

import (
	"context"
	"errors"
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
		return false, errors.New("no such user, please register")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(upass)); err != nil {
		return false, errors.New("password does not match")
	}

	return true, nil
}
