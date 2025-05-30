package persistent

import "context"

func (repo *PersistentRepo) GetWallet(ctx context.Context, uname string) (string, error) {
	if uname == "" {
		return "", errWalletEmptyUsername
	}

	type Wallet struct {
		ID string `db:"wallet_id"`
	}

	qGetWallet := "SELECT wallet_id from wallets WHERE customer_id = (SELECT id FROM customers c WHERE c.login = $1)"
	row := repo.QueryRowxContext(ctx, qGetWallet, uname)

	var wallet Wallet
	if err := row.StructScan(&wallet); err != nil {
		return "", errGenericDatabase
	}

	return wallet.ID, nil
}
