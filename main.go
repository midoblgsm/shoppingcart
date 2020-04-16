package main

import (
	"fmt"

	"github.com/midoblgsm/shoppingcart/cart"
	"github.com/midoblgsm/shoppingcart/resources"
)

func main() {
	cart := cart.NewCart()

	potato := resources.Item{
		ID:       "vg1",
		Name:     "Potato",
		Price:    2.5,
		Quantity: 2,
	}

	banana := resources.Item{
		ID:       "frt1",
		Name:     "Banana",
		Price:    1.5,
		Quantity: 3,
	}

	addReq := resources.AddItemRequest{Item: potato}
	addResp := cart.AddItem(addReq)
	if addResp.Error != nil {
		panic(addResp.Error.Error())
	}
	fmt.Printf("Cart %#v \n", cart)

	addResp = cart.AddItem(addReq)
	if addResp.Error != nil {
		panic(addResp.Error.Error())
	}
	fmt.Printf("Cart %#v \n", cart)

	addReq = resources.AddItemRequest{Item: banana}
	addResp = cart.AddItem(addReq)
	if addResp.Error != nil {
		panic(addResp.Error.Error())
	}
	fmt.Printf("Cart %#v \n", cart)

	removeRequest := resources.RemoveItemRequest{ItemID: "vg2", Quantity: 5}

	removeResp := cart.RemoveItem(removeRequest)
	if removeResp.Error != nil {
		panic(removeResp.Error.Error())
	}

	fmt.Printf("Cart %#v \n", cart)

	totalResp := cart.TotalCost()

	if totalResp.Error != nil {
		panic(totalResp.Error.Error())
	}

	fmt.Printf("Total Cost %f \n", totalResp.TotalCost)
}
