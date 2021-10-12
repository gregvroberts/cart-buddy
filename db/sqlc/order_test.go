package db

import (
	"context"
	"database/sql"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/*createRandomOrder Creates a random order in the database
@param t *testing.T The test object type
@return Order returns the generated order
*/
func createRandomOrder(t *testing.T) Order {

	usr := createRandomUser(t)

	arg := CreateOrderParams{
		OrderUserID:    usr.UserID,
		OrderAmount:    util.RandomFloat2(),
		OrderCity:      util.RandomString(30),
		OrderState:     util.RandomString(30),
		OrderPostal:    util.RandomPostal(),
		OrderCountry:   util.RandomString(30),
		OrderAddr1:     util.RandomAddress(),
		OrderPhone:     util.RandomPhone(),
		OrderShipping:  util.RandomFloat2(),
		OrderShipped:   false,
		OrderTrackCode: sql.NullString{String: util.RandomString(60), Valid: true},
	}

	ord1, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ord1)

	require.NotZero(t, ord1.OrderID)
	require.NotZero(t, ord1.CreatedAt)
	require.NotZero(t, ord1.UpdatedAt)
	require.NotZero(t, ord1.OrderDate)

	require.Equal(t, arg.OrderUserID, ord1.OrderUserID)
	require.Equal(t, arg.OrderAmount, ord1.OrderAmount)
	require.Equal(t, arg.OrderCity, ord1.OrderCity)
	require.Equal(t, arg.OrderState, ord1.OrderState)
	require.Equal(t, arg.OrderPostal, ord1.OrderPostal)
	require.Equal(t, arg.OrderCountry, ord1.OrderCountry)
	require.Equal(t, arg.OrderAddr1, ord1.OrderAddr1)
	require.Equal(t, arg.OrderAddr2, ord1.OrderAddr2)
	require.Equal(t, arg.OrderPhone, ord1.OrderPhone)
	require.Equal(t, arg.OrderShipping, ord1.OrderShipping)
	require.Equal(t, arg.OrderShipped, ord1.OrderShipped)
	require.Equal(t, arg.OrderTrackCode, ord1.OrderTrackCode)

	require.WithinDuration(t, ord1.CreatedAt, ord1.UpdatedAt, time.Second)
	require.WithinDuration(t, ord1.UpdatedAt, ord1.OrderDate, time.Second)
	require.WithinDuration(t, ord1.OrderDate, ord1.CreatedAt, time.Second)

	return ord1
}

/*TestQueries_CreateOrder Test the CreateOrder function
@param t *testing.T the test object type
@return NONE
*/
func TestQueries_CreateOrder(t *testing.T) {
	createRandomOrder(t)
}

/*TestQueries_GetOrder Test the GetOrder function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_GetOrder(t *testing.T) {
	ord1 := createRandomOrder(t)

	ord2, err := testQueries.GetOrder(context.Background(), ord1.OrderID)
	require.NoError(t, err)
	require.NotEmpty(t, ord2)

	require.Equal(t, ord1.OrderID, ord2.OrderID)
	require.Equal(t, ord1.OrderUserID, ord2.OrderUserID)
	require.Equal(t, ord1.OrderAmount, ord2.OrderAmount)
	require.Equal(t, ord1.OrderCity, ord2.OrderCity)
	require.Equal(t, ord1.OrderState, ord2.OrderState)
	require.Equal(t, ord1.OrderPostal, ord2.OrderPostal)
	require.Equal(t, ord1.OrderCountry, ord2.OrderCountry)
	require.Equal(t, ord1.OrderAddr1, ord2.OrderAddr1)
	require.Equal(t, ord1.OrderAddr2, ord2.OrderAddr2)
	require.Equal(t, ord1.OrderPhone, ord2.OrderPhone)
	require.Equal(t, ord1.OrderShipping, ord2.OrderShipping)
	require.Equal(t, ord1.OrderShipped, ord2.OrderShipped)
	require.Equal(t, ord1.OrderTrackCode, ord2.OrderTrackCode)

	require.NotZero(t, ord2.CreatedAt)
	require.NotZero(t, ord2.UpdatedAt)
	require.NotZero(t, ord2.OrderDate)

	require.WithinDuration(t, ord1.CreatedAt, ord2.CreatedAt, time.Second)
	require.WithinDuration(t, ord1.UpdatedAt, ord2.UpdatedAt, time.Second)
	require.WithinDuration(t, ord1.OrderDate, ord2.OrderDate, time.Second)

}

/*TestQueries_UpdateOrder Test the UpdateOrder function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_UpdateOrder(t *testing.T) {
	ord1 := createRandomOrder(t)
	user1 := createRandomUser(t)

	arg := UpdateOrderParams{
		OrderID:        ord1.OrderID,
		OrderUserID:    user1.UserID,
		OrderAmount:    util.RandomFloat2(),
		OrderCity:      util.RandomString(30),
		OrderState:     util.RandomString(30),
		OrderPostal:    util.RandomPostal(),
		OrderCountry:   util.RandomString(30),
		OrderAddr1:     util.RandomAddress(),
		OrderAddr2:     sql.NullString{String: util.RandomString(10), Valid: true},
		OrderPhone:     util.RandomPhone(),
		OrderShipping:  util.RandomFloat2(),
		OrderShipped:   true,
		OrderTrackCode: sql.NullString{String: util.RandomString(60), Valid: true},
	}

	ord2, err := testQueries.UpdateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ord2)

	require.Equal(t, ord1.OrderID, ord2.OrderID)
	require.NotEqual(t, ord1.OrderUserID, ord2.OrderUserID)
	require.NotEqual(t, ord1.OrderAmount, ord2.OrderAmount)
	require.NotEqual(t, ord1.OrderCity, ord2.OrderCity)
	require.NotEqual(t, ord1.OrderState, ord2.OrderState)
	require.NotEqual(t, ord1.OrderPostal, ord2.OrderPostal)
	require.NotEqual(t, ord1.OrderCountry, ord2.OrderCountry)
	require.NotEqual(t, ord1.OrderAddr1, ord2.OrderAddr1)
	require.NotEqual(t, ord1.OrderAddr2, ord2.OrderAddr2)
	require.NotEqual(t, ord1.OrderPhone, ord2.OrderPhone)
	require.NotEqual(t, ord1.OrderShipping, ord2.OrderShipping)
	require.NotEqual(t, ord1.OrderShipped, ord2.OrderShipped)
	require.NotEqual(t, ord1.OrderTrackCode, ord2.OrderTrackCode)
	require.Equal(t, ord1.CreatedAt, ord2.CreatedAt)

	require.NotZero(t, ord2.CreatedAt)
	require.NotZero(t, ord2.UpdatedAt)
	require.NotZero(t, ord2.OrderDate)

	require.NotEqual(t, ord1.UpdatedAt, ord2.UpdatedAt)
	require.NotEqual(t, ord1.OrderDate, ord2.OrderDate)
}

/*TestQueries_DeleteOrder tests the DeleteOrder function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_DeleteOrder(t *testing.T) {
	ord1 := createRandomOrder(t)

	err := testQueries.DeleteOrder(context.Background(), ord1.OrderID)
	require.NoError(t, err)

	ord2, err := testQueries.GetOrder(context.Background(), ord1.OrderID)
	require.Error(t, err)
	require.Empty(t, ord2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

/*TestQueries_ListOrders Tests the ListOrders function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_ListOrders(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}
	arg := ListOrdersParams{
		Limit:  5,
		Offset: 5,
	}

	orders, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, orders)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}
