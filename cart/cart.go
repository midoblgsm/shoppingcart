package cart

import (
	"fmt"
	"log"

	"github.com/midoblgsm/shoppingcart/resources"
)

type CartInterface interface {
	AddItem(resources.AddItemRequest) resources.AddItemResponse
	RemoveItem(resources.RemoveItemRequest) resources.RemoveItemResponse
	TotalItems() resources.TotalItemsResponse
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
	if (resources.Item{}) == request.Item {
		log.Printf("cannot-add-an-empty-item")
		return resources.AddItemResponse{Error: fmt.Errorf("cannot add an empty item")}
	}
	if request.Item.ID == "" {
		log.Printf("cannot-add-an-item-without-ID")
		return resources.AddItemResponse{Error: fmt.Errorf("cannot add an item without and ID")}
	}

	if request.Item.Price == 0 {
		log.Printf("cannot-add-an-item-without-price")
		return resources.AddItemResponse{Error: fmt.Errorf("cannot add an item without a price")}
	}

	if request.Item.Quantity == 0 {
		log.Printf("request-quantity-is-zero-defaulting-to-1")
		request.Item.Quantity = 1
	}

	existantItem, ok := c.Items[request.Item.ID]
	if !ok {
		c.Items[request.Item.ID] = request.Item
	}

	existantItem.Quantity = existantItem.Quantity + request.Item.Quantity
	c.Items[request.Item.ID] = existantItem

	c.Total = c.Total + (float32(request.Item.Quantity) * request.Item.Price)

	return resources.AddItemResponse{}
}

func (c *Cart) RemoveItem(request resources.RemoveItemRequest) resources.RemoveItemResponse {
	log.Println("entering-remove-item")
	defer log.Println("exiting-remove-item")

	return resources.RemoveItemResponse{}
}

func (c *Cart) TotalItems() resources.TotalItemsResponse {
	log.Println("entering-total-items")
	defer log.Println("exiting-total-items")
	return resources.TotalItemsResponse{TotalItems: len(c.Items)}
}

func (c *Cart) TotalCost() resources.TotalCostResponse {
	log.Println("entering-total-cost")
	defer log.Println("exiting-total-cost")
	return resources.TotalCostResponse{TotalCost: c.Total}
}
