package product

type Event interface {
	Apply(product *Product)
}

type ProductReceivedEvent struct {
	ID       string
	Quantity int
}

func (e ProductReceivedEvent) Apply(product *Product) {
	e.Quantity += e.Quantity
}

type ProductShippedEvent struct {
	ID       string
	Quantity int
}

func (e ProductShippedEvent) Apply(product *Product) {
	e.Quantity -= e.Quantity
}
