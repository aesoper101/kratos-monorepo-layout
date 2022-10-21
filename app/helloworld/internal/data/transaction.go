package data

import (
	"context"
	"fmt"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent"
)

func (data *Data) InTx(ctx context.Context, f func(context.Context) error) error {
	tx := ent.TxFromContext(ctx)
	if tx != nil {
		return f(ctx)
	}

	tx, err := data.db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	if err = f(ent.NewTxContext(ctx, tx)); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err2, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
