package cart

import (
	"fmt"
	"log"

	"github.com/midoblgsm/shoppingcart/resources"
)

type CartInterface interface {
	AddItem(resources.AddItemRequest) resources.AddItemResponse
	RemoveItem(resources.RemoveItemRequest) resources.RemoveItemResponse
	TotalCost() resources.TotalCostResponse
}

type Cart struct {
	Items map[string]resources.Item
	Total float32
}

func NewCart() CartInterface {
	return &Cart{
		Items: map[string]resources.Item{}}
}

func (c *Cart) AddItem(request resources.AddItemRequest) resources.AddItemResponse {
	log.Println("entering-add-item")
	defer log.Println("exiting-add-item")

	return resources.AddItemResponse{Error: fmt.Errorf("not implemented")}
}

func (c *Cart) RemoveItem(request resources.RemoveItemRequest) resources.RemoveItemResponse {
	log.Println("entering-remove-item")
	defer log.Println("exiting-remove-item")

	return resources.RemoveItemResponse{Error: fmt.Errorf("not implemented")}
}

func (c *Cart) TotalCost() resources.TotalCostResponse {
	return resources.TotalCostResponse{Error: fmt.Errorf("not implemented")}
	// return resources.TotalCostResponse{}
}
