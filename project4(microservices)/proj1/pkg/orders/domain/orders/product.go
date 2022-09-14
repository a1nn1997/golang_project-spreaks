package orders

import (
	"errors"
	"pkg/common/price"
)

type ProductID string

var ErrEmptyProductID = errors.New("empty product ID")

func NewProduct(id ProductID, name string, price price.Price) (Product, error){
	if len(id)==0{
		return Product{}, ErrEmptyProductID
	}
	return Product{id,name,price},nil
}

type Product struct {
	id ProductID
	name string
	price price.Price	
}

func (p Product) Name() string {
	return p.name
}


func (p Product) Price() price.Price {
	return p.price
}