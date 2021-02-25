package product

type State struct {
	QuantityOnHand int
}

type Product struct {
	id    string
	state State
}

func NewProduct(id string) Product {
	return Product{
		id:    id,
		state: State{},
	}
}
