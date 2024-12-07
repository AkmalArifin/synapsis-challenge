package models

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/synapsis-challenge/db"
)

type Cart struct {
	ID        int64      `json:"id"`
	UserID    null.Int64 `json:"user_id"`
	CreatedAt NullTime   `json:"created_at"`
	DeletedAt NullTime   `json:"deleted_at"`
}

type CartItem struct {
	ID        int64      `json:"id"`
	CartID    null.Int64 `json:"cart_id"`
	ProductID null.Int64 `json:"product_id"`
	Quantity  null.Int   `json:"quantity"`
	CreatedAt NullTime   `json:"created_at"`
	DeletedAt NullTime   `json:"deleted_at"`
}

func GetAllCarts() ([]Cart, error) {
	query := `SELECT id, user_id, created_at, deleted_at FROM carts WHERE deleted_at IS NULL`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var carts []Cart
	for rows.Next() {
		var cart Cart
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.DeletedAt)
		if err != nil {
			return nil, err
		}

		carts = append(carts, cart)
	}

	return carts, nil
}

func CreateCartByUser(userID int64) (Cart, error) {
	var cart Cart
	cart.UserID.SetValid(userID)
	err := cart.Save()
	if err != nil {
		return Cart{}, err
	}

	return cart, nil
}

func GetCartItemByID(cartItemID int64) (CartItem, error) {
	query := `
	SELECT id, cart_id, product_id, quantity, created_at, deleted_at 
	FROM cart_item 
	WHERE id = ? AND deleted_at IS NULL`

	var cartItem CartItem
	row := db.DB.QueryRow(query, cartItemID)
	err := row.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.CreatedAt, &cartItem.DeletedAt)
	if err != nil {
		return CartItem{}, err
	}
	return cartItem, nil
}

func GetCartByUser(userID int64) (Cart, error) {
	query := `SELECT id, user_id, created_at, deleted_at FROM carts WHERE user_id = ? AND deleted_at IS NULL`

	var cart Cart
	row := db.DB.QueryRow(query, userID)
	err := row.Scan(&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.DeletedAt)
	if err != nil {
		return Cart{}, err
	}

	return cart, nil
}

func GetCartItemsByCart(cartID int64) ([]CartItem, error) {
	query := `
	SELECT id, cart_id, product_id, quantity, created_at, deleted_at 
	FROM cart_item 
	WHERE cart_id = ? AND deleted_at IS NULL`

	rows, err := db.DB.Query(query, cartID)
	if err != nil {
		return nil, err
	}

	var cartItems []CartItem
	for rows.Next() {
		var cartItem CartItem
		err = rows.Scan(&cartItem.ID, &cartItem.CartID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.CreatedAt, &cartItem.DeletedAt)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}

func (c *Cart) Save() error {
	query := `
	INSERT INTO carts(user_id, created_at)
	VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	c.CreatedAt.SetValue(time.Now())
	result, err := stmt.Exec(c.UserID, c.CreatedAt)
	if err != nil {
		return err
	}

	c.ID, err = result.LastInsertId()

	return err
}

func (ci *CartItem) Save() error {
	query := `
	INSERT INTO cart_item(cart_id, product_id, quantity, created_at)
	VALUES(?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	ci.CreatedAt.SetValue(time.Now())
	result, err := stmt.Exec(ci.CartID, ci.ProductID, ci.Quantity, ci.CreatedAt)
	if err != nil {
		return err
	}

	ci.ID, err = result.LastInsertId()

	return err
}

func (ci *CartItem) Delete() error {
	query := `
	UPDATE cart_item
	SET deleted_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), ci.ID)

	return err
}
