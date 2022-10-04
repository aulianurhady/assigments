package routes

import (
	"github.com/aulianurhady/training/assignment2/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)

	// r.NoRoute(controllers.NoRoute)

	return r
}
