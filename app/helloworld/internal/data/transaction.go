package data

import (
	"context"
	"fmt"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent"
)

func (data *Data) InTx(ctx context.Context, f func(*ent.Tx) error) error {
	tx, err := data.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
		}
	}()
	if err := f(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
