package routes

import (
	"crud-go/connect"
	"crud-go/controllers/handlers"
	"github.com/gin-gonic/gin"
)

type perpus struct {
	connect connect.Connect
	router  *gin.Engine
}

func (p *perpus) Run() {
	handlers.NewAppRouter(p.connect, p.router).InitMainRoutes()
	if err := p.router.Run(p.connect.ApiServer()); err != nil {
		panic(err)
	}
}

func NewPerpusApp() *perpus {
	r := gin.Default()
	var appConnect = connect.NewConnect()
	return &perpus{connect: appConnect, router: r}
}
