package main

import (
	"github.com/midoblgsm/horizonDevops/shoppingcart/cart"
	"github.com/midoblgsm/horizonDevops/shoppingcart/resources"
)

func main() {
	cart := cart.NewCart()

	potato := resources.Item{
		ID:       "vg1",
		Name:     "Potato",
		Price:    1,
		Quantity: 2,
	}

	cart.AddItem(potato)
}
