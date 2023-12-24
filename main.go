package main

import (
	"github.com/joho/godotenv"
	"github.com/xairline/x-gpt/bootstrap"
	"github.com/xairline/x-gpt/utils"
	"go.uber.org/fx"
)

// @SecurityDefinitions.oauth2.accessCode
// @tokenUrl https://pluginlab.xairline.org/oauth/token
// @authorizationUrl https://pluginlab.xairline.org/oauth/authorize
// @scope all

// @BasePath	/apis
func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
