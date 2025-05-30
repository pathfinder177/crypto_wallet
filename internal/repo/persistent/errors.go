package persistent

import (
	"fmt"
)

var (
	errLoginNoUser          = fmt.Errorf("no such user, please register")
	errLoginNoMatchPassword = fmt.Errorf("password does not match")

	errRegBcryptGenFromPassword = fmt.Errorf("err: bcrypt GenerateFromPassword")

	RegViolateUniqueCustomers    = "customers_login_key"
	errRegViolateUniqueCustomers = fmt.Errorf("err: User already exists")

	RegViolatePKWallets    = "wallets_pkey"
	errRegViolatePKWallets = fmt.Errorf("err: Wallet already exists")
)
