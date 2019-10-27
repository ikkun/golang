package models

import (
	"database/sql"

	"../entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from golang_test")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Status:   status,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}
