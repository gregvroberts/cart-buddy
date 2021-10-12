package db

import (
	"context"
	"database/sql"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

/*createRandomProduct Inserts a random product into the database
@param t *testing.T The Test object type
@return NONE
*/
func createRandomProduct(t *testing.T) Product {
	pcat := createRandomProductCategory(t)

	arg := CreateProductParams{
		ProductSku:        util.RandomSKU(),
		ProductName:       util.RandomString(90),
		ProductPrice:      util.RandomFloat2(),
		ProductWeight:     util.RandomFloat2(),
		ProductCartDesc:   util.RandomString(200),
		ProductShortDesc:  util.RandomString(900),
		ProductLongDesc:   util.RandomString(10000),
		ProductThumb:      util.RandomImage(),
		ProductImage:      util.RandomImage(),
		ProductCategoryID: pcat.CategoryID,
		ProductLive:       util.RandomBool(),
		ProductLocation:   util.RandomString(10),
	}

	if i := rand.Intn(10); i%2 == 0 {
		arg.ProductStock = sql.NullFloat64{Valid: false}
		arg.ProductUnlimited = true
	} else {
		arg.ProductStock = sql.NullFloat64{Float64: util.RandomFloat2(), Valid: true}
		arg.ProductUnlimited = false
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.NotZero(t, product.ProductID)
	require.NotZero(t, product.CreatedAt)
	require.NotZero(t, product.UpdatedAt)
	require.NotZero(t, product.ProductUpdateDate)

	require.Equal(t, arg.ProductSku, product.ProductSku)
	require.Equal(t, arg.ProductName, product.ProductName)
	require.Equal(t, arg.ProductPrice, product.ProductPrice)
	require.Equal(t, arg.ProductWeight, product.ProductWeight)
	require.Equal(t, arg.ProductCartDesc, product.ProductCartDesc)
	require.Equal(t, arg.ProductShortDesc, product.ProductShortDesc)
	require.Equal(t, arg.ProductLongDesc, product.ProductLongDesc)
	require.Equal(t, arg.ProductThumb, product.ProductThumb)
	require.Equal(t, arg.ProductImage, product.ProductImage)
	require.Equal(t, arg.ProductCategoryID, product.ProductCategoryID)
	require.Equal(t, arg.ProductLive, product.ProductLive)
	require.Equal(t, arg.ProductLocation, product.ProductLocation)
	require.Equal(t, arg.ProductStock, product.ProductStock)
	require.Equal(t, arg.ProductUnlimited, product.ProductUnlimited)
	require.WithinDuration(t, product.CreatedAt, product.UpdatedAt, time.Second)
	require.WithinDuration(t, product.CreatedAt, product.ProductUpdateDate, time.Second)
	require.WithinDuration(t, product.UpdatedAt, product.ProductUpdateDate, time.Second)

	return product
}

/*TestQueries_CreateProduct Tests the CreateProduct function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_CreateProduct(t *testing.T) {
	createRandomProduct(t)
}

/*TestQueries_GetProduct Tests the GetProduct function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_GetProduct(t *testing.T) {
	p1 := createRandomProduct(t)
	p2, err := testQueries.GetProduct(context.Background(), p1.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, p2)

	require.Equal(t, p1.ProductID, p2.ProductID)
	require.Equal(t, p1.ProductSku, p2.ProductSku)
	require.Equal(t, p1.ProductName, p2.ProductName)
	require.Equal(t, p1.ProductPrice, p2.ProductPrice)
	require.Equal(t, p1.ProductWeight, p2.ProductWeight)
	require.Equal(t, p1.ProductCartDesc, p2.ProductCartDesc)
	require.Equal(t, p1.ProductShortDesc, p2.ProductShortDesc)
	require.Equal(t, p1.ProductLongDesc, p2.ProductLongDesc)
	require.Equal(t, p1.ProductThumb, p2.ProductThumb)
	require.Equal(t, p1.ProductImage, p2.ProductImage)
	require.Equal(t, p1.ProductCategoryID, p2.ProductCategoryID)
	require.Equal(t, p1.ProductLive, p2.ProductLive)
	require.Equal(t, p1.ProductLocation, p2.ProductLocation)
	require.Equal(t, p1.ProductStock, p2.ProductStock)
	require.Equal(t, p1.ProductUnlimited, p2.ProductUnlimited)
	require.Equal(t, p1.CreatedAt, p2.UpdatedAt)
	require.Equal(t, p1.UpdatedAt, p2.ProductUpdateDate)
	require.Equal(t, p1.ProductUpdateDate, p2.ProductUpdateDate)
}

/*TestQueries_DeleteProduct Tests the DeleteProduct function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_DeleteProduct(t *testing.T) {
	p1 := createRandomProduct(t)

	err := testQueries.DeleteProduct(context.Background(), p1.ProductID)
	require.NoError(t, err)

	p2, err := testQueries.GetProduct(context.Background(), p1.ProductID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, p2)
}

/*TestQueries_ListProducts Tests the ListProducts function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_ListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}
	arg := ListProductsParams{
		Limit:  5,
		Offset: 5,
	}

	products, err := testQueries.ListProducts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, products, 5)

	for _, product := range products {
		require.NotEmpty(t, product)
	}
}
