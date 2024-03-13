package main

import (
	"api-assignmet/controllers"
	"api-assignmet/lib"
	"api-assignmet/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	db , err := lib.InitDB()
	if err != nil {
		panic(err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderController := controllers.NewOrderController(orderRepository)

	ginEngine := gin.Default()

	ginEngine.POST("/orders", orderController.CreateOrder)
	ginEngine.GET("/orders", orderController.GetOrders)
	ginEngine.PUT("/orders/:orderId", orderController.UpdateOrder)
	ginEngine.DELETE("/orders/:orderId", orderController.DeleteOrder)

	err = ginEngine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}