package repository

import (
	"context"
	"database/sql"
)

type TxFn func(trx *sql.Tx) error

type BaseRepository struct {
	database *sql.DB
}

func (b *BaseRepository) WithTransaction(ctx context.Context, fn TxFn) (err error) {
	tx, err := b.database.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if tx == nil {
			return
		}

		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}

		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = rollbackErr
			}
		} else {
			commitErr := tx.Commit()
			if commitErr != nil {
				err = commitErr
			}
		}

	}()

	err = fn(tx)

	return err
}

func NewBaseRepository(database *sql.DB) BaseRepository {
	return BaseRepository{
		database: database,
	}
}
