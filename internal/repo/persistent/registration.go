package persistent

import (
	"context"
	"fmt"
	"main/internal/entity"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (repo *PersistentRepo) Create(ctx context.Context, reg entity.Registration) (bool, error) {
	uname := reg.Username
	upass := reg.Password
	uwallet := reg.Wallet

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(upass), bcrypt.DefaultCost)
	if err != nil {
		return false, errRegBcryptGenFromPassword
	}

	return repo._create(ctx, uname, string(passwordHash), uwallet)
}

func (repo *PersistentRepo) _create(ctx context.Context, uname, upass, uwallet string) (bool, error) {
	qRegisterUser := "INSERT INTO customers(created, id, login, password) VALUES (DEFAULT,DEFAULT, $1, $2) RETURNING id"
	qRegisterWallet := "INSERT INTO wallets(customer_id, wallet_id) VALUES ($1, $2)"

	tx, err := repo.BeginTxx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	var customer_id int
	err = tx.GetContext(ctx, &customer_id, qRegisterUser, uname, upass)
	if err != nil {
		if strings.Contains(fmt.Sprintln(err), RegViolateUniqueCustomers) {
			return false, errRegViolateUniqueCustomers
		}
		return false, err
	}

	_, err = tx.ExecContext(ctx, qRegisterWallet, customer_id, uwallet)
	if err != nil {
		if strings.Contains(fmt.Sprintln(err), RegViolatePKWallets) {
			return false, errRegViolatePKWallets
		}
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}
