package entities

import "fmt"

type Product struct {
	Id       int64
	Name     string
	Price    float64
	Quantity int64
	Status   bool
}

func (product Product) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nprice:%0.1f\nquantity: %d\nstatus: %t",
		product.Id, product.Name, product.Price, product.Quantity, product.Status)
}
