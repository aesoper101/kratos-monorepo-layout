package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(wire.Struct(new(Services), "*"), NewGreeterService)

type Services struct {
	*GreeterService
}
