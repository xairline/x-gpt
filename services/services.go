package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewDatarefService),
	fx.Provide(NewWebSocketService),
	fx.Provide(NewFlightLogsService),
)
