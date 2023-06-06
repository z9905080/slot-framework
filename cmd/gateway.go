package cmd

import (
	"context"
	"github.com/asim/go-micro/v3/registry"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"slot-framework/environment"
	slotGame "slot-framework/internal/gateway/domain/infra/slot_game"
	"slot-framework/internal/gateway/implement"
	"slot-framework/internal/gateway/interface/adapter"
	"slot-framework/internal/gateway/usecase"
	"slot-framework/pkg/logger"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// gameLogicCmd cmd
var gateCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use: "gateway",
		Run: func(cmd *cobra.Command, args []string) {
			gateFx()
		},
	}
	cmd.Flags().StringVar((*string)(&environment.ConfigPath), "config", "./environment/dev/config.json", "config file")

	return cmd
}

func gateFx() {

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

			// new logger
			logger.NewLogger,

			// new usecase
			usecase.NewUsecase,

			// new game logic client
			implement.NewGameLogicService,

			slotGame.NewGameManager,

			// new http server
			adapter.NewHTTPServer,
		),
		func() fx.Option {
			return fx.Provide(func() registry.Registry {
				return registry.DefaultRegistry
			})
		}(),
		fx.Invoke(gateExec),
	)

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}

	app.Run()
}

// handle app Lifecycle
func gateExec(lc fx.Lifecycle, f fx.Shutdowner, config environment.Config, reg registry.Registry, l logger.Logger, server http.Handler) error {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			decimal.DivisionPrecision = 20
			decimal.MarshalJSONWithoutQuotes = true

			l.InfoF("gateExec start")

			go func() {
				if err := http.ListenAndServe(":9102", server); err != nil {
					l.ErrorF("gateExec start error: %v", err)
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

			l.InfoF("gateExec stop")

			return nil
		},
	})
	return nil
}
