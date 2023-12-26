package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewMiscRoutes),
	fx.Provide(NewDatarefRoutes),
	fx.Provide(NewWebSocketRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	miscRoutes MiscRoutes,
	datarefRoutes DatarefRoutes,
	webSocketRoutes WebSocketRoutes,
) Routes {
	return Routes{
		miscRoutes,
		datarefRoutes,
		webSocketRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
