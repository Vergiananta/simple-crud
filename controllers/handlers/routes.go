package handlers

import (
	"crud-go/connect"
	"crud-go/manager"
	"crud-go/utils/response"
	"github.com/gin-gonic/gin"
)

type appRouter struct {
	connect connect.Connect
	router  *gin.Engine
}

func (r *appRouter) InitMainRoutes() {
	serviceManager := manager.NewUsecasManager(r.connect)
	NewCustomerController(r.router, response.NewJsonResponder(), serviceManager.CustomerUsecase()).InitRoutes()
}

func NewAppRouter(connect connect.Connect, app *gin.Engine) *appRouter {
	return &appRouter{
		connect,
		app,
	}
}
