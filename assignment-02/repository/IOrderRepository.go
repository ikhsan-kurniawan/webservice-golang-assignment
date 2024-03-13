package repository

import "api-assignmet/model"

type IOrderRepository interface {
	CreateOrder(newOrder model.Order) (model.Order, error)
	GetOrders() ([]model.Order, error)
	UpdateOrder(updatedOrder model.Order)(model.Order, error)
	DeleteOrder(orderId int)(string, error)
}