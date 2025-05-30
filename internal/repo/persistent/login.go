package persistent

import (
	"context"
	"main/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

func (repo *PersistentRepo) Read(ctx context.Context, login entity.Login) (bool, error) {
	uname := login.Username
	upass := login.Password

	type User struct {
		Login    string `db:"login"`
		Password string `db:"password"`
	}

	qLoginUser := "SELECT login, password FROM customers c WHERE c.login=$1"
	row := repo.QueryRowxContext(ctx, qLoginUser, uname)

	var user User
	err := row.StructScan(&user)
	if user.Login == "" || user.Password == "" {
		return false, errLoginNoUser
	}
	if err != nil {
		return false, errGenericDatabase
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(upass)); err != nil {
		return false, errLoginNoMatchPassword
	}

	return true, nil
}
