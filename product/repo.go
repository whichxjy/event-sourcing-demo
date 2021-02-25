package product

type ProductRepo struct {
	// product id => stream
	streams map[string]Stream
}

func (pr *ProductRepo) GetProduct(id string) Product {
	product := NewProduct(id)

	if stream, ok := pr.streams[id]; ok {
		for _, event := range stream.Events {
			product.ApplyEvent(event)
		}
	}

	return product
}
