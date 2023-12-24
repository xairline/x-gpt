package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewMiscRoutes),
	fx.Provide(NewDatarefRoutes),
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
) Routes {
	return Routes{
		miscRoutes,
		datarefRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
