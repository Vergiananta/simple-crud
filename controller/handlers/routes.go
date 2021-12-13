package handlers

import (
	"learn-go/connect"
	"learn-go/manager"
)

type appRouter struct {
	connect connect.Connect
}

func (r *appRouter) InitMainRoutes() {
	serviceManager := manager.NewUsecaseManager(r.connect)
	NewCustomerController(serviceManager.CustomerUsecase()).InitRoutes()
}

func NewAppRouter(connect connect.Connect) *appRouter {
	return &appRouter{
		connect,
	}
}
