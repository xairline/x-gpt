package controllers

import (
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewMiscController),
	fx.Provide(NewDatarefController),
	fx.Provide(NewWebSocketController),
)
