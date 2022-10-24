package main

import (
	"errors"
	"fmt"
	"github.com/aesoper101/kratos-monorepo-layout/app/helloworld/internal/conf"
	"github.com/aesoper101/kratos-utils/bootstrap"
	"github.com/aesoper101/kratos-utils/protobuf/types/confpb"
	"github.com/spf13/cobra"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Service go build -ldflags "-X Service.Id=x.y.z"
var (
	Service = bootstrap.NewServiceInfo(
		"kratos.admin",
		"1.0.0",
		"",
	)
)

func newRootCmd(args ...string) *cobra.Command {
	commandName := os.Args[0]
	if len(args) > 0 {
		commandName = args[0]
	}
	cmd := &cobra.Command{
		Use:   commandName,
		Short: "Run and manage",
		RunE: func(cmd *cobra.Command, args []string) error {
			var flags = bootstrap.GetFlags()
			return runApp(*flags)
		},
	}

	bootstrap.RegisterFlags(cmd.PersistentFlags())
	return cmd
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rc *confpb.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(bootstrap.NewRegistrarProvider(rc)),
	)
}

func runApp(cfg bootstrap.ConfigFlags) error {
	bc, cleanup, err := bootstrap.LoadConfig[conf.Bootstrap](cfg)
	if err != nil {
		return errors.New("load config failed")
	}
	defer cleanup()

	if bc.Tracer != nil {
		if err := bootstrap.NewTracerProvider(bc.Tracer, &Service); err != nil {
			return err
		}
	}

	logger := bootstrap.NewLoggerProvide(bc.Log, &Service)

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Registry, logger)
	if err != nil {
		return err
	}
	defer cleanup()

	err = bootstrap.InitOpenSergo(app, bc.Opensergo)
	if err != nil {
		return err
	}

	err = bootstrap.InitSentry(bc.Sentry)
	if err != nil {
		return err
	}

	return app.Run()
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
