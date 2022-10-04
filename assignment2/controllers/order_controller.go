package controllers

import (
	"strconv"

	"github.com/aulianurhady/training/assignment2/lib"
	"github.com/aulianurhady/training/assignment2/models"
	"github.com/aulianurhady/training/assignment2/repository"
	"github.com/aulianurhady/training/assignment2/transports"
	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Summary Create new Order
// @Description Create new Order
// @Param data body transports.Request true "Order data"
// @Success 200 {object} transports.Request "Order data"
// @Failure 400 {object} transports.Response
// @Router /orders [post]
// @Tags Orders
func CreateOrder(c *gin.Context) {
	db := lib.DB
	req := transports.Request{}
	var repoOrder repository.IOrderRepository
	var repoItem repository.IItemRepository

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	orderData := models.Orders{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
	}

	if err := repoOrder.CreateDataOrder(db, &orderData); err != nil {
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
		if err := repoItem.CreateDataItem(db, &itemData); err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
	}

	transports.SendResponse(c, req, nil)
}

// GetOrders godoc
// @Summary List of Orders
// @Description List of Orders
// @Success 200 {object} []transports.Request List of Orders
// @Failure 400 {object} transports.Response
// @Router /orders [get]
// @Tags Orders
func GetOrders(c *gin.Context) {
	db := lib.DB
	var responseData []transports.Response
	var repoOrder repository.IOrderRepository
	var repoItem repository.IItemRepository

	orderData, err := repoOrder.GetListOrder(db)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	for _, v := range orderData {
		itemData, err := repoItem.GetListItemByID(db, v.OrderID)
		if err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
		res := lib.BuildResponseData(v, itemData)

		responseData = append(responseData, res)
	}

	transports.SendResponse(c, responseData, nil)
}

// UpdateOrder godoc
// @Summary Update Orders by id
// @Description Update Orders by id
// @Param data body transports.Request true "Order data"
// @Success 200 {object} transports.Request "Order data"
// @Failure 400 {object} transports.Response
// @Router /orders [put]
// @Tags Orders
func UpdateOrder(c *gin.Context) {
	db := lib.DB
	req := transports.Request{}
	var repoOrder repository.IOrderRepository
	var repoItem repository.IItemRepository

	if err := c.BindJSON(&req); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	orderID, _ := strconv.Atoi(c.Param("orderId"))

	orderData, err := repoOrder.GetDataOrderByID(db, orderID)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	updatedOrderData := models.Orders{
		OrderID:      orderData.OrderID,
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
	}

	if err := repoOrder.UpdateDataOrder(db, &updatedOrderData); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	for _, v := range req.CustomerItems {
		itemData, err := repoItem.GetDataItemByID(db, v.LineItemID)
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

		if err := repoItem.UpdateDataItem(db, &updatedItemData); err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
	}

	transports.SendResponse(c, req, nil)
}

// RemoveOrder godoc
// @Summary Delete Orders by id
// @Description Delete Orders by id
// @Param id path string true "Orders ID"
// @Success 200 {object} transports.Response
// @Failure 400 {object} transports.Response
// @Router /orders/{id} [delete]
// @Tags Orders
func RemoveOrder(c *gin.Context) {
	db := lib.DB
	var repoOrder repository.IOrderRepository
	var repoItem repository.IItemRepository

	orderID, _ := strconv.Atoi(c.Param("orderId"))

	orderData, err := repoOrder.GetDataOrderByID(db, orderID)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	if err := repoOrder.DeleteDataOrder(db, &orderData); err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	itemList, err := repoItem.GetListItemByID(db, orderID)
	if err != nil {
		transports.SendResponse(c, nil, err)
		return
	}

	for _, v := range itemList {
		if err := repoItem.DeleteDataItem(db, &v); err != nil {
			transports.SendResponse(c, nil, err)
			return
		}
	}

	transports.SendResponse(c, gin.H{
		"orderID": orderID,
		"status":  "deleted",
	}, nil)
}
