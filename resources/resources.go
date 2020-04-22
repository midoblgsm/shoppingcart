package resources

type Config struct {
	Port int
}

type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type AddItemRequest struct {
	Item Item `json:"item"`
}

type AddItemResponse struct {
	Error error `json:"error"`
}

type RemoveItemRequest struct {
	ItemID   string `json:"itemid"`
	Quantity int    `json:"quantity"`
}

type RemoveItemResponse struct {
	Error error `json:"error"`
}

type TotalItemsResponse struct {
	TotalItems int   `json:"totalitems"`
	Error      error `json:"error"`
}

type TotalCostResponse struct {
	TotalCost float32 `json:"totalcost"`
	Error     error   `json:"error"`
}
