package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type storage struct {
	db *sqlx.DB
}

func New(dsn string) (*storage, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &storage{
		db: db,
	}, nil
}

func (s *storage) GetCustomer(ctx context.Context, id int) {

}
