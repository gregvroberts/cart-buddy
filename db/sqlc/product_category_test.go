package db

import (
	"context"
	"database/sql"
	"github.com/gregvroberts/cart-buddy/util"
	"github.com/stretchr/testify/require"
	"testing"
)

/*createRandomProductCategory Creates a random product category in the database
@param t *testing.T The test object type
@return ProductCategory The generated ProductCategory
*/
func createRandomProductCategory(t *testing.T) ProductCategory {
	categoryName := util.RandomString(15)

	pcat, err := testQueries.CreateProductCategory(context.Background(), categoryName)
	require.NoError(t, err)
	require.NotEmpty(t, pcat)
	require.Equal(t, categoryName, pcat.CategoryName)
	require.NotZero(t, pcat.CategoryID)

	return pcat
}

/*TestQueries_CreateProductCategory Tests the CreateProductCategory function
@param t *testing.T The test object type
@return none
*/
func TestQueries_CreateProductCategory(t *testing.T) {
	createRandomProductCategory(t)
}

/*TestQueries_GetProductCategory Tests the GetProductCategory function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_GetProductCategory(t *testing.T) {
	pcat1 := createRandomProductCategory(t)

	pcat2, err := testQueries.GetProductCategory(context.Background(), pcat1.CategoryID)
	require.NoError(t, err)
	require.NotEmpty(t, pcat2)
	require.Equal(t, pcat1.CategoryID, pcat2.CategoryID)
	require.Equal(t, pcat2.CategoryName, pcat2.CategoryName)
}

/*TestQueries_UpdateProductCategory Tests the UpdateProductCategory function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_UpdateProductCategory(t *testing.T) {
	pcat1 := createRandomProductCategory(t)

	arg := UpdateProductCategoryParams{
		CategoryID:   pcat1.CategoryID,
		CategoryName: util.RandomString(30),
	}

	pcat2, err := testQueries.UpdateProductCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pcat2)
	require.Equal(t, pcat1.CategoryID, pcat2.CategoryID)
	require.NotEqual(t, pcat1.CategoryName, pcat2.CategoryName)
}

/*TestQueries_DeleteProductCategory Tests the DeleteProductCategory function
@param t *testing.T The test object type
@return NONE
*/
func TestQueries_DeleteProductCategory(t *testing.T) {
	pcat1 := createRandomProductCategory(t)
	err := testQueries.DeleteProductCategory(context.Background(), pcat1.CategoryID)
	require.NoError(t, err)

	pcat2, err := testQueries.GetProductCategory(context.Background(), pcat1.CategoryID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, pcat2)
}

/*TestQueries_ListProductCategories Tests the ListProductCategories function
@param t *testing.T The Test object type
@return NONE
*/
func TestQueries_ListProductCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProductCategory(t)
	}

	arg := ListProductCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	pcats, err := testQueries.ListProductCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, pcats, 5)

	for _, pcat := range pcats {
		require.NotEmpty(t, pcat)
	}
}
