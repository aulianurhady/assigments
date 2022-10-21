package routes

import (
	"github.com/aulianurhady/training/mygram/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	// Users
	r.POST("/users/register", controllers.UserRegister)
	// r.POST("/users/login", controllers.GetOrders)
	// r.PUT("/users", controllers.UpdateOrder)
	// r.DELETE("/users", controllers.RemoveOrder)

	return r
}
