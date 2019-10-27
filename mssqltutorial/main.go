package main

import (
	"fmt"

	"./config"
	"./models"
)

func Demo1() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
			}
		}
	}
}
