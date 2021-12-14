package handlers

import (
	"crud-go/connect"
	"crud-go/manager"
)

type appRouter struct {
	connect connect.Connect
}

func (r *appRouter) InitMainRoutes() {
	serviceManager := manager.NewUsecasManager(r.connect)
	NewCustomerController(serviceManager.CustomerUsecase()).InitRoutes()
}

func NewAppRouter(connect connect.Connect) *appRouter {
	return &appRouter{connect}
}
