package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	groupName    string
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *OrderHandler {
	return &OrderHandler{"api/order", orderService}
}

func (handler *OrderHandler) RegisterEndPoints(r *gin.Engine) {
	orderGroup := r.Group(handler.groupName)

	orderGroup.POST("", handler.CreateOrder)
	orderGroup.GET("", handler.GetOrders)
	orderGroup.GET("/:id", handler.GetOrder)
	orderGroup.POST("/:id", handler.UpdateOrder)
	orderGroup.DELETE("/:id", handler.DeleteOrder)
}

func (handler *OrderHandler) CreateOrder(ctx *gin.Context) {

	orderData := NewInputCreateOrder()
	err := ctx.BindJSON(&orderData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to bind order data"})
		return
	}
	newOrder, err := handler.orderService.CreateOrder(orderData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to create new order"})
return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "order created", "data": newOrder})

}

func (handler *OrderHandler) GetOrders(ctx *gin.Context) {
	allOrders, err := handler.orderService.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to get orders data"})
		return
	}
	// ctx.JSON(http.StatusOK, gin.H{"msg": "orders data fetched successfully", "data": allOrders})
	ctx.JSON(http.StatusOK, allOrders)
}

func (handler *OrderHandler) GetOrder(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid order id"})
		return
	}
	order, err := handler.orderService.GetOrder(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to get order data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "order data fetched successfully", "data": order})
}

func (handler *OrderHandler) UpdateOrder(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid order id"})
		return
	}
	orderData := NewInputUpdateOrder()
	err := ctx.BindJSON(&orderData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "failed to bind orderdata"})
		return
	}
	updatedOrderData, err := handler.orderService.UpdateOrder(id, orderData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to update order data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "order data updated successfully", "data": updatedOrderData})
}

func (handler *OrderHandler) DeleteOrder(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "invalid order id"})
		return
	}
	err := handler.orderService.DeleteOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to delete order"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "order deleted"})
}
