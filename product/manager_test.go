package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductManager(t *testing.T) {
	pm := GetProductManager()

	{
		product := pm.Get("a")
		assert.Equal(t, product.state.QuantityOnHand, 0)
		product.ReceiveProduct(10)
		assert.Equal(t, product.state.QuantityOnHand, 10)
		pm.Save(product)
	}
	{
		product := pm.Get("a")
		assert.Equal(t, product.state.QuantityOnHand, 10)
	}

	{
		product := pm.Get("b")
		product.ReceiveProduct(10)
		product.ShipProduct(5)
		product.ReceiveProduct(20)
		product.ShipProduct(6)
		product.ReceiveProduct(30)
		product.ShipProduct(7)
		product.ReceiveProduct(40)
		product.ShipProduct(8)
		product.ReceiveProduct(50)
		product.ShipProduct(9)
		pm.Save(product)
	}
	{
		product := pm.Get("b")
		assert.Equal(t, product.state.QuantityOnHand, 115)
	}
}
