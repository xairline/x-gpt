package main

import (
	"github.com/joho/godotenv"
	"github.com/xairline/x-gpt/bootstrap"
	"github.com/xairline/x-gpt/utils"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
