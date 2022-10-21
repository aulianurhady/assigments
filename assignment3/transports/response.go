package transports

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func SendResponse(c *gin.Context, data Response, err error) {
	c.Header("Content-Type", "text/html")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"status": data,
	})
}
