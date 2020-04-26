package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midoblgsm/shoppingcart/resources"
)

type CartServer struct {
	ctx     context.Context
	handler CartHandler
	config  resources.Config
}

//NewDataStreamServer creates a new instance of the DataStreamServer
func NewCartServer(c context.Context, h CartHandler, conf resources.Config) *CartServer {
	return &CartServer{ctx: c, handler: h, config: conf}
}

//InitializeHandler prepares the handlers to be used for the rest queries
func (c *CartServer) InitializeHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/v1/item", c.handler.AddItem()).Methods("POST")
	router.HandleFunc("/v1/item", c.handler.RemoveItem()).Methods("DELETE")
	router.HandleFunc("/v1/cart/totalCost", c.handler.TotalCost()).Methods("GET")
	router.HandleFunc("/v1/cart/items", c.handler.GetItems()).Methods("GET")
	return router
}

//Start starts the http server
func (c *CartServer) Start() error {
	router := c.InitializeHandler()
	http.Handle("/", router)
	c.printStartMsg()

	return http.ListenAndServe(fmt.Sprintf(":%d", c.config.Port), nil)
}
func (c *CartServer) printStartMsg() {
	log.Printf(fmt.Sprintf("Starting shoppingcart on port %d ....", c.config.Port))
	log.Printf("CTL-C to exit/stop shoppingcart service")
}
