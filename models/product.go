package models

import (
	"github.com/guregu/null/v5"
	"github.com/synapsis-challenge/db"
)

type Product struct {
	ID          int64       `json:"id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Price       null.Int    `json:"price"`
	CategoryID  null.Int64  `json:"category_id"`
	CreatedAt   NullTime    `json:"created_at"`
	DeletedAt   NullTime    `json:"deleted_at"`
}

func GetAllProducts() ([]Product, error) {
	query := `
	SELECT id, name, description, price, category_id, created_at, deleted_at
	FROM products
	WHERE deleted_at IS NULL
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.CreatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProductByID(productID int64) (Product, error) {
	query := `
	SELECT id, name, description, price, category_id, created_at, deleted_at
	FROM products
	WHERE id = ? AND deleted_at IS NULL`

	var product Product
	row := db.DB.QueryRow(query, productID)
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.CreatedAt, &product.DeletedAt)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func GetProductsByCategory(categoryID int64) ([]Product, error) {
	query := `
	SELECT id, name, description, price, category_id, created_at, deleted_at
	FROM products
	WHERE category_id = ? AND deleted_at IS NULL
	`
	rows, err := db.DB.Query(query, categoryID)
	if err != nil {
		return nil, err
	}

	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.CreatedAt, &product.DeletedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
