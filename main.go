package main

import (
	"context"
	"fmt"
	"log"

	"github.com/midoblgsm/shoppingcart/cart"
	"github.com/midoblgsm/shoppingcart/resources"
	"github.com/midoblgsm/shoppingcart/server"
)

func main() {
	fmt.Println("coorect code")
	ctx := context.Background()
	config := resources.Config{Port: 7778}
	cart := cart.NewCart()

	handler := server.NewCartHandler(ctx, cart)
	server := server.NewCartServer(ctx, handler, config)

	log.Printf("starting-server %#v", server.Start())
}
