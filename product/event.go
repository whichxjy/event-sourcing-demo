package product

type ProductReceivedEvent struct {
	ID       string
	Quantity int
}

type ProductShippedEvent struct {
	ID       string
	Quantity int
}
