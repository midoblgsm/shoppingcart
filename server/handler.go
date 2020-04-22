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

//CreateDataStream creates a handlerFunc for crating a new PullPush DataStream
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
