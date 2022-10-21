package routes

import (
	"github.com/aulianurhady/training/mygram/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	// Users
	r.POST("/users/register", controllers.UserRegister)
	r.POST("/users/login", controllers.UserLogin)
	r.PUT("/users", controllers.UserUpdate)
	r.DELETE("/users", controllers.UserDelete)

	// Photos
	r.POST("/photos", controllers.PhotoInsert)
	r.GET("/photos", controllers.GetListPhotos)
	r.PUT("/photos/:photoID", controllers.PhotoUpdate)
	r.DELETE("/photos/:photoID", controllers.PhotoDelete)

	return r
}
