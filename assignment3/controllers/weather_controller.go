package controllers

import (
	"github.com/aulianurhady/training/assignment3/lib"
	"github.com/aulianurhady/training/assignment3/transports"
	"github.com/gin-gonic/gin"
)

func GetWeatherStatus(c *gin.Context) {
	var responseData = transports.Response{
		Water: lib.GetRandomInt(),
		Wind:  lib.GetRandomInt(),
	}

	transports.SendResponse(c, responseData, nil)
}
