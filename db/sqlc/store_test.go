package db

import (
	"context"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAddToCartTx(t *testing.T) {
	store := NewStore(testDB)

	// run n concurrent AddToCart transactions
	n := 2

	errs := make(chan error)
	results := make(chan AddToCartTxResult)

	prod1 := createRandomProduct(t)
	ord1 := createRandomOrder(t)

	for i := 0; i < n; i++ {
		go func() {

			result, err := store.AddToOrderTx(context.Background(), AddToCartTxParams{
				ProductID: prod1.ProductID,
				OrderID:   ord1.OrderID,
				Quantity:  util.RandomInt(1, 5),
			})
			errs <- err
			results <- result
		}()
	}

	// Check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check Product
		product := result.Product
		require.NotEmpty(t, product)

		// Check order
		order := result.Order
		require.NotEmpty(t, order)

		// Check OrderDetail
		orderDetail := result.OrderDetail
		require.NotEmpty(t, orderDetail)
		require.Equal(t, product.ProductID, orderDetail.DetailProductID)
		require.Equal(t, product.ProductName, orderDetail.DetailProductName)
		require.Equal(t, product.ProductSku, orderDetail.DetailSku)
		require.Equal(t, product.ProductPrice, orderDetail.DetailUnitPrice)
		require.Equal(t, result.Quantity, orderDetail.DetailQuantity)
		require.WithinDuration(t, orderDetail.CreatedAt, orderDetail.UpdatedAt, time.Second)
	}

}
