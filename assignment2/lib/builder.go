package lib

import (
	"github.com/aulianurhady/training/assignment2/models"
	"github.com/aulianurhady/training/assignment2/transports"
)

func BuildResponseData(orderData models.Orders, itemData []models.Items) transports.Response {
	var itemRespData transports.ItemResponse
	responseData := transports.Response{
		OrderID:      orderData.OrderID,
		CustomerName: orderData.CustomerName,
	}

	for _, v := range itemData {

		itemRespData.ItemID = v.ItemID
		itemRespData.ItemCode = v.ItemCode
		itemRespData.Description = v.Description
		itemRespData.Quantity = v.Quantity

		responseData.CustomerItems = append(responseData.CustomerItems, itemRespData)
	}

	return responseData
}
