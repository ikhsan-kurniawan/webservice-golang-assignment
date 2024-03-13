package controllers

import (
	"api-assignmet/model"
	"api-assignmet/repository"
	"api-assignmet/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderRepository repository.IOrderRepository
}

func NewOrderController(orderRepository repository.IOrderRepository) *orderController {
	return &orderController{
		orderRepository: orderRepository,
	}
}

func (oc *orderController) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	createdOrder, err := oc.orderRepository.CreateOrder(newOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true,createdOrder, ""))
}


func (oc *orderController) GetOrders(ctx *gin.Context) {
	orders, err := oc.orderRepository.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, orders, ""))
}

func (oc *orderController) UpdateOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	var updatedOrder model.Order

	err = ctx.ShouldBindJSON(&updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	
	updatedOrder.OrderId = orderID
	
	updatedOrder, err = oc.orderRepository.UpdateOrder(updatedOrder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updatedOrder, ""))
}


func (oc *orderController) DeleteOrder(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.CreateResponse(false, nil, err.Error()))
		return
	}

	pesan, err := oc.orderRepository.DeleteOrder(orderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, pesan, ""))
}