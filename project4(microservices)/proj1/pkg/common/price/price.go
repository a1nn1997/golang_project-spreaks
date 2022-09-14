package price

import (
	"errors"
)

var(
	ErrPriceTooLow      = errors.New("price must be greater than 0")
	ErrInvalidCurrency  = errors.New(" invalid currency")
)
type price struct{
	cents 				uint64
	currency			string
}

func NewPrice( cents uint64, currency string)(Price, error){
	if cents<=0{
		return Price{},ErrPriceTooLow
	}
	if len(currency)!=3{
		return Price{},ErrInvalidCurrency
	}
	return Price{cents, currency},nil

}

func NewPriceP( cents uint64, currency string)(Price){
	p,err := NewPrice(cents, currency)
	if err!=nil{
		panic(err)
	}
	return p
}