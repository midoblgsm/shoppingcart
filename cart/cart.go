package cart

import (
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

	existantItem, ok := c.Items[request.Item.ID]
	if !ok {
		c.Items[request.Item.ID] = request.Item
	} else {
		existantItem.Quantity = existantItem.Quantity + request.Item.Quantity
		c.Items[request.Item.ID] = existantItem
	}
	c.Total = c.Total + (float32(request.Item.Quantity) * request.Item.Price)
	return resources.AddItemResponse{}
}

func (c *Cart) RemoveItem(request resources.RemoveItemRequest) resources.RemoveItemResponse {
	log.Println("entering-remove-item")
	defer log.Println("exiting-remove-item")
	if request.Quantity < 1 {
		return resources.RemoveItemResponse{}
	}

	existantItem, ok := c.Items[request.ItemID]
	if !ok {
		return resources.RemoveItemResponse{}
	}

	if existantItem.Quantity < request.Quantity {
		delete(c.Items, request.ItemID)
		c.Total = c.Total - (float32(existantItem.Quantity) * existantItem.Price)
	} else {
		existantItem.Quantity = existantItem.Quantity - request.Quantity
		c.Items[request.ItemID] = existantItem
		c.Total = c.Total - (float32(request.Quantity) * existantItem.Price)
	}

	return resources.RemoveItemResponse{}
}
func (c *Cart) TotalCost() resources.TotalCostResponse {
	return resources.TotalCostResponse{
		TotalCost: c.Total}
}
