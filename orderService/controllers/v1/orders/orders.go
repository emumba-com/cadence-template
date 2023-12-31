package orders

import (
	"github.com/gin-gonic/gin"
	"orderService/clients/cadenceClient"
	"orderService/handlers/v1/orders"
)

func registerRoutes(server *orders.Server) {
	orderRoutes := server.RouterGroup.Group("orders")
	{
		orderRoutes.POST("/", server.PlaceAnOrder)
	}

}

func CreateNewServer(router *gin.Engine, rg *gin.RouterGroup, cc cadenceClient.CadStore) {
	server := &orders.Server{
		Router:        router,
		RouterGroup:   rg,
		CadenceClient: cc,
	}

	registerRoutes(server)
}
