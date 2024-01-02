package bootstrap

import (
	"context"
	"github.com/xairline/x-gpt/controllers"
	"github.com/xairline/x-gpt/routes"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	utils.Module,
	controllers.Module,
	routes.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler utils.RequestHandler,
	routes routes.Routes,
	env utils.Env,
	logger utils.Logger,
) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {

			go func() {
				routes.Setup()
				host := "0.0.0.0"
				handler.Gin.Run(host + ":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Application")
			return nil
		},
	})
}
