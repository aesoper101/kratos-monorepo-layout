package server

import (
	"context"
	"github.com/aesoper101/go-utils/validatorx"
	v1 "github.com/aesoper101/kratos-monorepo-layout/api/helloworld/v1"

	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/service"
	"github.com/aesoper101/kratos-utils/pkg/middleware/translator"
	"github.com/aesoper101/kratos-utils/pkg/middleware/validate"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	ut "github.com/go-playground/universal-translator"
	jwtv4 "github.com/golang-jwt/jwt/v4"

	"github.com/aesoper101/kratos-utils/pkg/middleware/metrics"
	"github.com/aesoper101/kratos-utils/pkg/middleware/requestid"
	"github.com/aesoper101/kratos-utils/pkg/network"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	sentrykratos "github.com/go-kratos/sentry"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, services *service.Services, trans *ut.UniversalTranslator, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
				// do someting
				return nil
			})),
			sentrykratos.Server(),
			tracing.Server(),
			logging.Server(logger),
			ratelimit.Server(),
			metadata.Server(),
			requestid.Server(),
			jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(c.Grpc.AuthKey.Value), nil
			}),
			translator.Translate(translator.WithUniversalTranslator(trans)),
			validate.Validator(validate.WithRegisterDefaultTranslations(validatorx.RegisterDefaultTranslations)),
			metrics.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	if tlsCfg := c.Grpc.GetTls(); tlsCfg != nil {
		opts = append(opts, grpc.TLSConfig(network.InitTLSConfig(tlsCfg)))
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, services.GreeterService)
	return srv
}
