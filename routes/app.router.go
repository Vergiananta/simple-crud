package routes

import (
	"learn-go/connect"
	"learn-go/controller/handlers"
)

type perpus struct {
	connect connect.Connect
}

func (p *perpus) Run() {
	handlers.NewAppRouter(p.connect).InitMainRoutes()
}

func NewPerpusApp() *perpus {
	var appConnect = connect.NewConnect()
	return &perpus{connect: appConnect}
}
