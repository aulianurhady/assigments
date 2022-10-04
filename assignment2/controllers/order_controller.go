package controllers

import (
	"fmt"
	"strconv"

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
		transports.SendResponse(c, nil, err)
		return
	}

	orderData := models.Orders{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
	}

	if err := repository.CreateDataOrder(db, &orderData); err != nil {
		transports.SendResponse(c, nil, err)
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
			transports.SendResponse(c, nil, err)
			return
		}
	}

	transports.SendResponse(c, req, nil)
}

func GetOrders(c *gin.Context) {
	db := lib.DB
	var responseData []transports.Response

	orderData, err := repository.GetListOrder(db)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	for _, v := range orderData {
		itemData, err := repository.GetListItemByID(db, v.OrderID)
		if err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
		res := lib.BuildResponseData(v, itemData)

		responseData = append(responseData, res)
	}

	transports.SendResponse(c, responseData, nil)
}

func UpdateOrder(c *gin.Context) {
	db := lib.DB
	req := transports.Request{}

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	orderID, _ := strconv.Atoi(c.Param("orderId"))

	orderData, err := repository.GetDataOrderByID(db, orderID)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	updatedOrderData := models.Orders{
		OrderID:      orderData.OrderID,
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
	}

	if err := repository.UpdateDataOrder(db, &updatedOrderData); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	for _, v := range req.CustomerItems {
		itemData, err := repository.GetDataItemByID(db, v.LineItemID)
		if err != nil {
			transports.SendResponse(c, nil, err)
			return
		}

		updatedItemData := models.Items{
			ItemID:      itemData.ItemID,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
		}

		if err := repository.UpdateDataItem(db, &updatedItemData); err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
	}

	transports.SendResponse(c, req, nil)
}

func RemoveOrder(c *gin.Context) {
	db := lib.DB

	orderID, _ := strconv.Atoi(c.Param("orderId"))

	orderData, err := repository.GetDataOrderByID(db, orderID)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	if err := repository.DeleteDataOrder(db, &orderData); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	itemList, err := repository.GetListItemByID(db, orderID)
	if err != nil {
		fmt.Println("MASUK SINI?")
		transports.SendResponse(c, nil, err)
		return
	}

	for _, v := range itemList {
		if err := repository.DeleteDataItem(db, &v); err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
	}

	transports.SendResponse(c, gin.H{
		"orderID": orderID,
		"status":  "deleted",
	}, nil)
}
