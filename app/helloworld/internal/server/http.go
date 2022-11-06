package server

import (
	"context"
	"github.com/aesoper101/go-utils/validatorx"
	v1 "github.com/aesoper101/kratos-monorepo-layout/api/helloworld/v1"

	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/service"
	"github.com/aesoper101/kratos-utils/pkg/encoder"
	"github.com/aesoper101/kratos-utils/pkg/middleware/metrics"
	"github.com/aesoper101/kratos-utils/pkg/middleware/realip"
	"github.com/aesoper101/kratos-utils/pkg/middleware/requestid"
	"github.com/aesoper101/kratos-utils/pkg/middleware/translator"
	"github.com/aesoper101/kratos-utils/pkg/middleware/validate"
	"github.com/aesoper101/kratos-utils/pkg/network"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	sentrykratos "github.com/go-kratos/sentry"
	ut "github.com/go-playground/universal-translator"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, services *service.Services, trans *ut.UniversalTranslator, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
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
			translator.Translate(translator.WithUniversalTranslator(trans)),
			validate.Validator(validate.WithRegisterDefaultTranslations(validatorx.RegisterDefaultTranslations)),
			metrics.Server(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST"}),
			),
			realip.Handler,
		),
		http.ResponseEncoder(encoder.ApiEncodeResponse()),
		http.ErrorEncoder(encoder.ApiErrorEncoder()),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	if tlsCfg := c.Http.GetTls(); tlsCfg != nil {
		opts = append(opts, http.TLSConfig(network.InitTLSConfig(tlsCfg)))
	}

	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, services.GreeterService)

	network.RegisterPprof(srv, c.Http.GetPprof())
	network.RegisterMetrics(srv, c.Http.GetMetrics())
	network.RegisterSwagger(srv, c.Http.GetSwagger())
	return srv
}
