package server

import (
	"context"
	"log"
	"net/http"

	"github.com/midoblgsm/shoppingcart/cart"
	"github.com/midoblgsm/shoppingcart/resources"
	"github.com/midoblgsm/shoppingcart/utils"
)

type CartHandler struct {
	ctx  context.Context
	cart cart.CartInterface
}

func NewCartHandler(ct context.Context, c cart.CartInterface) CartHandler {
	return CartHandler{ctx: ct, cart: c}
}

//AddItem creates a handlerFunc for adding an Item
func (c *CartHandler) AddItem() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("entering-AddItem")
		defer log.Printf("exiting-AddItem")

		addItemRequest := resources.AddItemRequest{}
		err := utils.UnmarshalDataFromRequest(req, &addItemRequest)
		if err != nil {
			log.Printf("error-unmarshalling-%s", err.Error())
			utils.WriteResponse(w, 409, &resources.AddItemResponse{Error: err})
			return
		}

		resp := c.cart.AddItem(addItemRequest)
		if resp.Error != nil {
			log.Printf("error-writing-response-%s", resp.Error.Error())
			utils.WriteResponse(w, 409, &resp)
			return
		}
		log.Printf("cart %#v", c.cart)
		utils.WriteResponse(w, http.StatusOK, resp)
	}
}

//RemoveItem creates a handlerFunc for removing an Item
func (c *CartHandler) RemoveItem() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("entering-RemoveItem")
		defer log.Printf("exiting-RemoveItem")

		removeItemRequest := resources.RemoveItemRequest{}
		err := utils.UnmarshalDataFromRequest(req, &removeItemRequest)
		if err != nil {
			log.Printf("error-unmarshalling-data-%s", err.Error())
			utils.WriteResponse(w, 409, &resources.RemoveItemResponse{Error: err})
			return
		}

		resp := c.cart.RemoveItem(removeItemRequest)
		if resp.Error != nil {
			log.Printf("error-removing-item")
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, resp)
	}
}

func (c *CartHandler) TotalCost() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("entering-TotalCost")
		defer log.Printf("exiting-TotalCost")

		resp := c.cart.TotalCost()
		if resp.Error != nil {
			log.Printf("error-getting-totalcost")
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, resp)
	}
}

func (c *CartHandler) GetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("entering-getItems")
		defer log.Printf("exiting-getItems")

		resp := c.cart.GetItems()
		if resp.Error != nil {
			log.Printf("error-getting-items")
			utils.WriteResponse(w, 409, &resp)
			return
		}
		utils.WriteResponse(w, http.StatusOK, resp)
	}
}
