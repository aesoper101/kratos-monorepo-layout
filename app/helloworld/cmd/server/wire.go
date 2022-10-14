//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/biz"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/server"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/service"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *confpb.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
