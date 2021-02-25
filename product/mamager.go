package product

type ProductManager struct {
	// product id => stream
	streams map[string]*Stream
}

func (pm *ProductManager) Get(id string) Product {
	product := NewProduct(id)

	if stream, ok := pm.streams[id]; ok {
		for _, event := range stream.Events {
			product.ApplyEvent(event)
		}
	}

	return product
}

func (pm *ProductManager) Save(product Product) {
	if stream, ok := pm.streams[product.id]; !ok || stream == nil {
		pm.streams[product.id] = &Stream{}
	}

	stream := pm.streams[product.id]
	newEvents := product.GetUncommittedEvents()
	stream.Events = append(stream.Events, newEvents...)
	product.EventsCommitted()
}
