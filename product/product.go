package product

type State struct {
	QuantityOnHand int
}

type Product struct {
	id                string
	state             State
	uncommittedEvents []Event
	allEvents         []Event
}

func NewProduct(id string) Product {
	return Product{
		id:    id,
		state: State{},
	}
}

func (p *Product) ApplyEvent(event Event) {
	event.Apply(p)
	p.allEvents = append(p.allEvents, event)
}

func (p *Product) addUncommittedEvent(event Event) {
	p.ApplyEvent(event)
	p.uncommittedEvents = append(p.uncommittedEvents, event)
}

func (p *Product) ReceiveProduct(quantity int) {
	p.addUncommittedEvent(ProductReceivedEvent{
		ID:       p.id,
		Quantity: quantity,
	})
}

func (p *Product) ShipProduct(quantity int) {
	if quantity > p.state.QuantityOnHand {
		panic("Not enough product to ship")
	}

	p.addUncommittedEvent(ProductShippedEvent{
		ID:       p.id,
		Quantity: quantity,
	})
}

func (p *Product) GetUncommittedEvents() []Event {
	return append([]Event(nil), p.uncommittedEvents...)
}

func (p *Product) GetAllEvents() []Event {
	return append([]Event(nil), p.allEvents...)
}

func (p *Product) EventsCommitted() {
	p.uncommittedEvents = nil
}
