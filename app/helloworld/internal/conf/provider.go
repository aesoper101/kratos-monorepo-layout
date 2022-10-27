package conf

import (
	"github.com/aesoper101/kratos-utils/bootstrap"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/google/wire"
)

var (
	ProviderSet = wire.NewSet(
		NewConfigLoader,
		GetConfig,
		GetTracerConfig,
		GetLogConfig,
		GetRegistryConfig,
		GetOpenSergoConfig,
		GetSentryConfig,
		GetServerConfig,
		GetDataBaseConfig,
		GetHttpConfig,
		GetGRPCConfig,
	)
)

func NewConfigLoader(cfg bootstrap.ConfigFlags) (*bootstrap.ConfigLoader[Bootstrap], func(), error) {
	return bootstrap.NewConfigLoader[Bootstrap](cfg)
}

func GetConfig(loader *bootstrap.ConfigLoader[Bootstrap]) *Bootstrap {
	return loader.GetConfig()
}

func GetTracerConfig(cfg *Bootstrap) *confpb.Tracer {
	return cfg.Tracer
}

func GetLogConfig(cfg *Bootstrap) *confpb.Log {
	return cfg.Log
}

func GetRegistryConfig(cfg *Bootstrap) *confpb.Registry {
	return cfg.Registry
}

func GetOpenSergoConfig(cfg *Bootstrap) *confpb.OpenSergo {
	return cfg.Opensergo
}

func GetSentryConfig(cfg *Bootstrap) *confpb.Sentry {
	return cfg.Sentry
}

func GetServerConfig(cfg *Bootstrap) *Server {
	return cfg.Server
}

func GetDataBaseConfig(cfg *Bootstrap) *Data {
	return cfg.Data
}

func GetHttpConfig(cfg *Server) *confpb.HTTP {
	return cfg.Http
}

func GetGRPCConfig(cfg *Server) *confpb.GRPC {
	return cfg.Grpc
}
