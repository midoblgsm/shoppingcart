package resources

type Item struct {
	ID       string
	Name     string
	Price    float32
	Quantity int
}

type AddItemRequest struct {
	Item Item
}

type AddItemResponse struct {
	Error error
}

type RemoveItemRequest struct {
	ItemID   string
	Quantity int
}

type RemoveItemResponse struct {
	Error error
}

type TotalCostResponse struct {
	TotalCost float32
	Error     error
}
