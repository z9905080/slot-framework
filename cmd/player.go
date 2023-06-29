package cmd

import (
	"context"
	"github.com/asim/go-micro/v3/registry"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
	"slot-framework/environment"
	"slot-framework/internal/player/domain/service"
	"slot-framework/internal/player/interface/adapter"
	"slot-framework/internal/player/interface/handler"
	"slot-framework/internal/player/usecase"
	"slot-framework/pkg/logger"
)

// playerCmd cmd
var playerCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use: "player",
		Run: func(cmd *cobra.Command, args []string) {
			playerFx()
		},
	}
	cmd.Flags().StringVar((*string)(&environment.ConfigPath), "config", "./environment/dev/config.json", "config file")

	return cmd
}

func playerFx() {
	// di
	app := fx.New(
		fx.NopLogger,
		fx.Supply(
			environment.ConfigPath,
		),
		fx.Provide(
			context.Background,

			// new env
			environment.New,

			adapter.NewMicroServer,

			handler.NewPlayerHandler,

			usecase.NewPlayerUsecase,

			service.NewPlayerService,

			// new logger
			logger.NewLogger,
		),

		func() fx.Option {
			return fx.Provide(func() registry.Registry {
				return registry.DefaultRegistry
			})
		}(),
		fx.Invoke(playerExec),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func playerExec(lc fx.Lifecycle, f fx.Shutdowner, l logger.Logger, grpcServer *adapter.GrpcServer, reg registry.Registry) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			l.InfoF("playerExec start")

			go func() {
				if err := grpcServer.Run(); err != nil {
					l.ErrorF("grpcServer.Run err: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			f.Shutdown()

			serviceList, err := reg.GetService(reg.String())
			if err != nil {
				return err
			}

			for _, srv := range serviceList {
				if deRegisterErr := reg.Deregister(srv); deRegisterErr != nil {
					return deRegisterErr
				}
			}

			return nil
		},
	})
	return nil
}
