package main

import (
	"context"

	"github.com/NpoolPlatform/file-gateway/api"

	"github.com/NpoolPlatform/file-gateway/pkg/feeder"
	"github.com/NpoolPlatform/file-gateway/pkg/service"

	action "github.com/NpoolPlatform/go-service-framework/pkg/action"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/oss"
	ossconst "github.com/NpoolPlatform/go-service-framework/pkg/oss/const"

	apicli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		defer feeder.Shutdown(c.Context)
		defer service.Shutdown(c.Context)

		return action.Run(
			c.Context,
			run,
			rpcRegister,
			rpcGatewayRegister,
			watch,
		)
	},
}

func run(ctx context.Context) error {
	if err := oss.Init(ossconst.ImageStoreKey, "image_bucket"); err != nil {
		return err
	}
	return nil
}

func shutdown(ctx context.Context) {
	<-ctx.Done()
	logger.Sugar().Infow(
		"Watch",
		"State", "Done",
		"Error", ctx.Err(),
	)
}

func _watch(ctx context.Context, cancel context.CancelFunc, w func(ctx context.Context)) {
	defer func() {
		if err := recover(); err != nil {
			logger.Sugar().Errorw(
				"Watch",
				"State", "Panic",
				"Error", err,
			)
			cancel()
		}
	}()
	w(ctx)
}

func watch(ctx context.Context, cancel context.CancelFunc) error {
	go shutdown(ctx)
	go _watch(ctx, cancel, service.Watch)
	go _watch(ctx, cancel, feeder.Watch)
	return nil
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	apicli.RegisterGRPC(server)

	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return err
	}

	_ = apicli.Register(mux)

	return nil
}
