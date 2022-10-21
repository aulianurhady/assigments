package routes

import (
	"github.com/aulianurhady/training/assignment3/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/weather/status", controllers.GetWeatherStatus)

	return r
}
