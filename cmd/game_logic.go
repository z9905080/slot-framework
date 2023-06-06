package cmd

import (
	"context"
	"github.com/asim/go-micro/v3/registry"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"log"
	"slot-framework/environment"
	"slot-framework/internal/game_logic/domain/service"
	"slot-framework/internal/game_logic/implement"
	"slot-framework/internal/game_logic/interface/adapter"
	"slot-framework/internal/game_logic/interface/handler"
	"slot-framework/internal/game_logic/usecase"
	"slot-framework/pkg/logger"
)

// gameLogicCmd cmd
var gameLogicCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use: "game_logic",
		Run: func(cmd *cobra.Command, args []string) {
			gameLogicFx()
		},
	}
	cmd.Flags().StringVar((*string)(&environment.ConfigPath), "config", "./environment/dev/config.json", "config file")

	return cmd
}

func gameLogicFx() {
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

			handler.NewGameLogicHandler,

			usecase.NewGameLogicUsecase,

			implement.NewGameLogicServiceServer,

			service.NewGameService,

			// new logger
			logger.NewLogger,
		),

		func() fx.Option {
			return fx.Provide(func() registry.Registry {
				return registry.DefaultRegistry
			})
		}(),
		fx.Invoke(gameLogicExec),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

func gameLogicExec(lc fx.Lifecycle, f fx.Shutdowner, l logger.Logger, grpcServer *adapter.GrpcServer, reg registry.Registry) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			l.InfoF("gateExec start")

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
