package controllers

import (
	"github.com/aulianurhady/training/assignment2/lib"
	"github.com/aulianurhady/training/assignment2/models"
	"github.com/aulianurhady/training/assignment2/repository"
	"github.com/aulianurhady/training/assignment2/transports"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := lib.DB
	req := transports.Request{}

	if err := c.BindJSON(&req); err != nil {
		return
	}

	orderData := models.Orders{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
	}

	if err := repository.CreateDataOrder(db, &orderData); err != nil {
		return
	}

	for _, v := range req.CustomerItems {
		itemData := models.Items{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     orderData.OrderID,
		}
		if err := repository.CreateDataItem(db, &itemData); err != nil {
			return
		}
	}

	transports.SendResponse(c, req)
}

func GetOrders(c *gin.Context) {
	db := lib.DB
	var (
		// orderData    []models.Orders
		// err          error
		responseData []transports.Response
	)

	// responseData := make([]transports.Response)

	orderData, err := repository.GetDataOrder(db)
	if err != nil {
		return
	}

	for _, v := range orderData {
		itemData, err := repository.GetDataItemByID(db, v.OrderID)
		if err != nil {
			return
		}
		res := lib.BuildResponseData(v, itemData)

		responseData = append(responseData, res)
	}

	transports.SendResponse(c, responseData)
}
