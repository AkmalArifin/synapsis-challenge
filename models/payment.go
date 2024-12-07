package models

import (
	"time"

	"github.com/guregu/null/v5"
	"github.com/synapsis-challenge/db"
)

type Payment struct {
	ID        int64       `json:"id"`
	OrderID   null.Int64  `json:"order_id"`
	Amount    null.Int    `json:"amount"`
	Provider  null.String `json:"provider"`
	Status    null.String `json:"status"`
	CreatedAt NullTime    `json:"created_at"`
	DeletedAt NullTime    `json:"deleted_at"`
}

func GetAllPayments() ([]Payment, error) {
	query := `
	SELECT id, order_id, amount, provider, status, created_at, deleted_at
	FROM payments
	WHERE deleted_at IS NULL
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var payments []Payment
	for rows.Next() {
		var payment Payment
		err = rows.Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Provider, &payment.Status, &payment.CreatedAt, &payment.DeletedAt)
		if err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func (p *Payment) Save() error {
	query := `
	INSERT INTO payments(order_id, amount, provider, status, created_at)
	VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	p.CreatedAt.SetValue(time.Now())
	result, err := stmt.Exec(p.OrderID, p.Amount, p.Provider, p.Status, p.CreatedAt)
	if err != nil {
		return err
	}

	p.ID, err = result.LastInsertId()

	return err
}
