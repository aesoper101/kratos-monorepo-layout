package biz

import (
	"context"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data/ent"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterLogic)

type Transaction interface {
	InTx(context.Context, func(ctx *ent.Tx) error) error
}
