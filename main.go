package main

import (
	"github.com/joho/godotenv"
	"github.com/xairline/x-gpt/bootstrap"
	"github.com/xairline/x-gpt/utils"
	"go.uber.org/fx"
)

// @SecurityDefinitions.oauth2.accessCode
// @tokenUrl https://auth.xairline.org/oidc/token
// @authorizationUrl https://auth.xairline.org/oidc/auth
// @scope email

// @BasePath	/apis
func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
