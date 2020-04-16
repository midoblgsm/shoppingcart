package cart

import (
	"fmt"

	"github.com/midoblgsm/horizonDevops/shoppingcart/resources"
)

type CartInterface interface {
	AddItem(resources.Item) error
	RemoveItem(string, int) error
	TotalAmount() (float32, error)
	TotalUniqueItems() (int, error)
	TotalUnits() (int, error)
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

func (c *Cart) TotalAmount() (float32, error) {
	return 0, nil
}
func (c *Cart) TotalUniqueItems() (int, error) {
	return 0, nil
}
func (c *Cart) TotalUnits() (int, error) {
	return 0, nil
}
