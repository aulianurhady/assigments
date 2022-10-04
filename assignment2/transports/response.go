package transports

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	OrderID       int            `json:"order_id"`
	CustomerName  string         `json:"customer_name"`
	CustomerItems []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	ItemID      int    `json:"item_id"`
}

func SendResponse(c *gin.Context, data interface{}, err error) {
	c.Header("Content-Type", "application/json")

	if data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": map[string]interface{}{
				"id":      "ERR-01",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}
