package repository

import (
	"api-assignmet/model"
	"database/sql"
	"errors"
	"fmt"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository{
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) CreateOrder(newOrder model.Order) (model.Order, error) {
	tx, err := or.db.Begin()
	if err != nil {
		return model.Order{}, err
	}


	queryOrders := "insert into orders (customer_name, ordered_at) values($1, $2) returning *"

	row := tx.QueryRow(queryOrders, newOrder.CustomerName, newOrder.OrderedAt)
	err = row.Scan(&newOrder.OrderId, &newOrder.CustomerName, &newOrder.OrderedAt)
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	for k, v := range newOrder.Items {
		queryItems := "insert into items (item_code, description, quantity, order_id) values ($1, $2, $3, $4) returning item_id, order_id"
		err := tx.QueryRow(queryItems, v.ItemCode, v.Description, v.Quantity, newOrder.OrderId).Scan(&newOrder.Items[k].ItemId, &newOrder.Items[k].OrderId)
		if err != nil {
			tx.Rollback()
			return model.Order{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}
	
	return newOrder, nil
}

func (or *orderRepository) GetOrders() ([]model.Order, error) {
	var orders = []model.Order{}

	query := `
	select
		o.order_id,
		o.customer_name,
		o.ordered_at,
		i.item_id,
		i.item_code,
		i.description,
		i.quantity,
		i.order_id
	from orders o
	JOIN items i ON o.order_id = i.order_id
	ORDER BY o.order_id
	`

	rows, err := or.db.Query(query)
	if err != nil {
		return orders, err
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		var item model.Item

		err := rows.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
		if err != nil {
			continue
		}

		var isExist bool
		for i := range orders {
			if orders[i].OrderId == order.OrderId {
				orders[i].Items = append(orders[i].Items, item)
				isExist = true
				break
			}
		}

		if !isExist {
			orders = append(orders, model.Order{
				OrderId:      order.OrderId,
				CustomerName: order.CustomerName,
				OrderedAt:    order.OrderedAt,
				Items:        []model.Item{item},
			})
		}

	}

	return orders, nil
}

func (or *orderRepository) UpdateOrder(updatedOrder model.Order) (model.Order, error) {
	tx, err := or.db.Begin()
	if err != nil {
		return model.Order{}, err
	}

	fmt.Printf("%+v", updatedOrder)

	queryOrders := "UPDATE orders SET customer_name = $1, ordered_at = $2 WHERE order_id = $3"
	_, err = tx.Exec(queryOrders, updatedOrder.CustomerName, updatedOrder.OrderedAt, updatedOrder.OrderId)
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	for _, v := range updatedOrder.Items {
		queryItems := "UPDATE items SET item_code = $1, description = $2, quantity = $3 WHERE item_id = $4 AND order_id = $5"
		_, err := tx.Exec(queryItems, v.ItemCode, v.Description, v.Quantity, v.ItemId, updatedOrder.OrderId)
		if err != nil {
			tx.Rollback()
			return model.Order{}, err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	return updatedOrder, nil
}

func (or *orderRepository) DeleteOrder(orderId int) (string, error) {
    query := "SELECT COUNT(*) FROM orders WHERE order_id = $1"
    var hitung int
    err := or.db.QueryRow(query, orderId).Scan(&hitung)
    if err != nil {
        return "", err
    }

	if hitung == 0 {
		return "", errors.New("data tidak ditemukan")
	}
	
    tx, err := or.db.Begin()
    if err != nil {
        return "", err
    }
    queryItem := "DELETE FROM items WHERE order_id = $1"
    _, err = tx.Exec(queryItem, orderId)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    queryOrder := "DELETE FROM orders WHERE order_id = $1"
    _, err = tx.Exec(queryOrder, orderId)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    err = tx.Commit()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    return "Data berhasil dihapus", nil
}