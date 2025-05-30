package persistent

import (
	"fmt"
)

var (
	errGenericDatabase = fmt.Errorf("something went wrong on server side")

	errLoginNoUser          = fmt.Errorf("no such user, please register")
	errLoginNoMatchPassword = fmt.Errorf("password does not match")

	errRegBcryptGenFromPassword = fmt.Errorf("err: bcrypt GenerateFromPassword")

	RegViolateUniqueCustomers    = "customers_login_key"
	errRegViolateUniqueCustomers = fmt.Errorf("err: User already exists")

	RegViolatePKWallets    = "wallets_pkey"
	errRegViolatePKWallets = fmt.Errorf("err: Wallet already exists")

	errWalletEmptyUsername = fmt.Errorf("err: Username is empty as get the wallet")
)
