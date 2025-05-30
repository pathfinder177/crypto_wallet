package persistent

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PersistentRepo struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *PersistentRepo {
	return &PersistentRepo{db}
}

func SetConfig(db *sqlx.DB) {
	//FIXME move to config
	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(15)
	db.SetConnMaxIdleTime(time.Hour * 5)
}
