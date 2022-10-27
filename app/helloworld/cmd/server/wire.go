//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/biz"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/data"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/i18n"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/server"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/service"
	"github.com/aesoper101/kratos-utils/bootstrap"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(cfg bootstrap.ConfigFlags, srvInfo *bootstrap.ServiceInfo) (*bootstrap.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		i18n.ProviderSet,
		conf.ProviderSet,
		bootstrap.ProviderSet,
	))
}
