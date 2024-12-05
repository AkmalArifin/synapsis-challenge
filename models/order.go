package models

import (
	"fmt"
	"time"

	"github.com/guregu/null/v5"
	"github.com/synapsis-challenge/db"
)

type Order struct {
	ID        int64      `json:"id"`
	UserID    null.Int64 `json:"user_id"`
	Amount    null.Int   `json:"amount"`
	CreatedAt NullTime   `json:"created_at"`
	DeletedAt NullTime   `json:"deleted_at"`
}

type OrderItem struct {
	ID        int64      `json:"id"`
	OrderID   null.Int64 `json:"order_id"`
	ProductID null.Int64 `json:"product_id"`
	Quantity  null.Int   `json:"quantity"`
	CreatedAt NullTime   `json:"created_at"`
	DeletedAt NullTime   `json:"deleted_at"`
}

func CreateOrderFromCart(cart Cart) (Order, []OrderItem, error) {
	cartItems, err := GetCartItemsByCart(cart.ID)
	if err != nil {
		return Order{}, nil, err
	}

	fmt.Println(cartItems)

	var amount int = 0
	var orderItems []OrderItem
	for i := 0; i < len(cartItems); i++ {
		var orderItem OrderItem
		orderItem.ProductID.SetValid(cartItems[i].ProductID.ValueOrZero())
		orderItem.Quantity.SetValid(cartItems[i].Quantity.ValueOrZero())
		orderItem.CreatedAt.SetValue(time.Now())
		orderItems = append(orderItems, orderItem)

		// Get the total mount of order
		product, err := GetProductByID(orderItem.ProductID.Int64)
		if err != nil {
			return Order{}, nil, err
		}
		amount += int(orderItem.Quantity.ValueOrZero() * product.Price.ValueOrZero())

		// delete cart item because of checkout
		err = cartItems[i].Delete()
		if err != nil {
			return Order{}, nil, err
		}
	}

	var order Order
	order.UserID.SetValid(cart.UserID.ValueOrZero())
	order.Amount.SetValid(int64(amount))
	err = order.Save()
	if err != nil {
		return Order{}, nil, err
	}

	for i := 0; i < len(orderItems); i++ {
		orderItems[i].OrderID.SetValid(order.ID)
		err = orderItems[i].Save()
		if err != nil {
			return Order{}, nil, err
		}
	}

	return order, orderItems, nil
}

func (o *Order) Save() error {
	query := `
	INSERT INTO orders(user_id, amount, created_at)
	VALUES(?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	o.CreatedAt.SetValue(time.Now())
	result, err := stmt.Exec(o.UserID, o.Amount, o.CreatedAt)
	if err != nil {
		return err
	}

	o.ID, err = result.LastInsertId()

	return err
}

func (oi *OrderItem) Save() error {
	query := `
	INSERT INTO order_item(order_id, product_id, quantity, created_at)
	VALUES(?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	oi.CreatedAt.SetValue(time.Now())
	result, err := stmt.Exec(oi.OrderID, oi.ProductID, oi.Quantity, oi.CreatedAt)
	if err != nil {
		return err
	}

	oi.ID, err = result.LastInsertId()

	return err
}
