package cart

import (
	"fmt"

	"github.com/midoblgsm/shoppingcart/resources"
)

type CartInterface interface {
	AddItem(resources.Item) error
	RemoveItem(string, int) error
	TotalCost() (float32, error)
}

type Cart struct {
	Items map[string]resources.Item
}

func NewCart() CartInterface {
	return &Cart{
		Items: map[string]resources.Item{}}
}

func (c *Cart) AddItem(item resources.Item) error {
	fmt.Println("Not implemented")
	return nil
}

func (c *Cart) RemoveItem(itemID string, qt int) error {
	return nil
}

func (c *Cart) TotalCost() (float32, error) {
	return 0, nil
}
