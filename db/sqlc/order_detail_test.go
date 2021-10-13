package db

import (
	"context"
	"database/sql"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"testing"
)

/*createRandomOrderDetail Generates a random OrderDetail record
@param t *testing.T The test object type
@return NONE
*/
func createRandomOrderDetail(t *testing.T) OrderDetail {
	prod1 := createRandomProduct(t)
	ord1 := createRandomOrder(t)

	arg := CreateOrderDetailParams{
		DetailOrderID:     ord1.OrderID,
		DetailProductID:   prod1.ProductID,
		DetailProductName: prod1.ProductName,
		DetailUnitPrice:   prod1.ProductPrice,
		DetailSku:         prod1.ProductSku,
		DetailQuantity:    util.RandomInt(1, 100),
	}

	orderDetail1, err := testQueries.CreateOrderDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail1)

	require.NotZero(t, orderDetail1.DetailID)
	require.NotZero(t, orderDetail1.CreatedAt)
	require.NotZero(t, orderDetail1.UpdatedAt)

	require.Equal(t, arg.DetailOrderID, orderDetail1.DetailOrderID)
	require.Equal(t, arg.DetailProductID, orderDetail1.DetailProductID)
	require.Equal(t, arg.DetailProductName, orderDetail1.DetailProductName)
	require.Equal(t, arg.DetailUnitPrice, orderDetail1.DetailUnitPrice)
	require.Equal(t, arg.DetailSku, orderDetail1.DetailSku)
	require.Equal(t, arg.DetailQuantity, orderDetail1.DetailQuantity)

	return orderDetail1
}

/*TestQueries_CreateOrderDetail Tests the CreateOrderDetail function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_CreateOrderDetail(t *testing.T) {
	createRandomOrderDetail(t)
}

/*TestQueries_GetOrderDetail Tests the GetOrderDetail function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_GetOrderDetail(t *testing.T) {
	orderDetail1 := createRandomOrderDetail(t)

	orderDetail2, err := testQueries.GetOrderDetail(context.Background(), orderDetail1.DetailID)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetail2)

	require.Equal(t, orderDetail1.DetailID, orderDetail2.DetailID)
	require.Equal(t, orderDetail1.DetailOrderID, orderDetail2.DetailOrderID)
	require.Equal(t, orderDetail1.DetailProductID, orderDetail2.DetailProductID)
	require.Equal(t, orderDetail1.DetailProductName, orderDetail2.DetailProductName)
	require.Equal(t, orderDetail1.DetailUnitPrice, orderDetail2.DetailUnitPrice)
	require.Equal(t, orderDetail1.DetailSku, orderDetail2.DetailSku)
	require.Equal(t, orderDetail1.DetailQuantity, orderDetail2.DetailQuantity)
	require.Equal(t, orderDetail1.CreatedAt, orderDetail2.CreatedAt)
	require.Equal(t, orderDetail1.UpdatedAt, orderDetail2.UpdatedAt)
}

/*TestQueries_UpdateOrderDetail Tests the UpdateOrderDetail function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_UpdateOrderDetail(t *testing.T) {

}

/*TestQueries_DeleteOrderDetail Tests the DeleteOrderDetail function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_DeleteOrderDetail(t *testing.T) {

	orderDetail1 := createRandomOrderDetail(t)

	err := testQueries.DeleteOrderDetail(context.Background(), orderDetail1.DetailID)
	require.NoError(t, err)

	orderDetail2, err := testQueries.GetOrderDetail(context.Background(), orderDetail1.DetailOrderID)
	require.Error(t, err)
	require.Empty(t, orderDetail2)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

/*TestQueries_ListOrderDetail Tests the ListOrderDetail function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_ListOrderDetails(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrderDetail(t)
	}

	arg := ListOrderDetailsParams{
		Limit:  5,
		Offset: 5,
	}

	orderDetails, err := testQueries.ListOrderDetails(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orderDetails)
	require.Len(t, orderDetails, 5)

	for _, orderDetail := range orderDetails {
		require.NotEmpty(t, orderDetail)
	}
}
