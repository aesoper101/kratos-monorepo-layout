package main

import (
	"fmt"
	"github.com/aesoper101/kratos-utils/bootstrap"
	"github.com/spf13/cobra"
	"os"
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

func runApp(cfg bootstrap.ConfigFlags) error {
	app, cleanup, err := wireApp(cfg, &Service)
	if err != nil {
		return err
	}
	defer cleanup()

	return app.Run()
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
