package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {

	// Declare variables
	var product1 Product
	var product2 Product
	var product3_1 Product
	var product4_1 Product
	var product5_1 Product
	var product6_2 Product
	var product7_1 Product
	var product8 Product
	var err error

	// Delete test data when testing ends
	defer func() {
		deleteTestProduct(testQueries, product1.ProductID)
		deleteTestProduct(testQueries, product2.ProductID)
		deleteTestProduct(testQueries, product3_1.ProductID)
		deleteTestProduct(testQueries, product4_1.ProductID)
		deleteTestProduct(testQueries, product5_1.ProductID)
		deleteTestProduct(testQueries, product6_2.ProductID)
		deleteTestProduct(testQueries, product7_1.ProductID)
		deleteTestProduct(testQueries, product8.ProductID)
	}()

	// Create some test data in the DB
	product1, err = insertRandomProduct(testQueries)
	require.NoError(t, err)
	product2, err = insertRandomProduct(testQueries)
	require.NoError(t, err)
	product3_1, err = insertRandomProductWithOwner(testQueries, product1.ShopOwnerName)
	require.NoError(t, err)
	product4_1, err = insertRandomProductWithOwner(testQueries, product1.ShopOwnerName)
	require.NoError(t, err)
	product5_1, err = insertRandomProductWithOwner(testQueries, product1.ShopOwnerName)
	require.NoError(t, err)
	product6_2, err = insertRandomProductWithOwner(testQueries, product2.ShopOwnerName)
	require.NoError(t, err)
	product7_1, err = insertRandomProductWithOwner(testQueries, product1.ShopOwnerName)
	require.NoError(t, err)
	product8, err = insertRandomProduct(testQueries)
	require.NoError(t, err)

	// Call GetProductCount
	count, err := testQueries.GetProductCount(context.Background())
	// Check GetProductCount
	require.NoError(t, err)
	require.NotEmpty(t, count)
	require.Equal(t, int64(8), count)

	// Call GetProductCountByOwner
	count, err = testQueries.GetProductCountByOwner(context.Background(), product1.ShopOwnerName)
	// Check GetProductCountByOwner
	require.NoError(t, err)
	require.NotEmpty(t, count)
	require.Equal(t, int64(5), count)

	// Call GetProductList for OK
	productList, err := testQueries.GetProductList(context.Background(), GetProductListParams{
		Page:     2,
		Pagesize: 3,
	})
	// Check GetProductList for OK
	require.NoError(t, err)
	require.NotEmpty(t, productList)
	require.Equal(t, productList, []GetProductListRow{
		createProductListRow(product4_1),
		createProductListRow(product5_1),
		createProductListRow(product6_2),
	},
	)

	// Call GetProductListByOwner for OK
	productListByOwner, err := testQueries.GetProductListByOwner(context.Background(), GetProductListByOwnerParams{
		ShopOwnerName: product1.ShopOwnerName,
		Page:          1,
		Pagesize:      4,
	})
	// Check GetProductListByOwner for OK
	require.NoError(t, err)
	require.NotEmpty(t, productListByOwner)
	require.Equal(t, productListByOwner, []GetProductListByOwnerRow{
		createProductListByOwnerRow(product1),
		createProductListByOwnerRow(product3_1),
		createProductListByOwnerRow(product4_1),
		createProductListByOwnerRow(product5_1),
	},
	)

	// Call GetProduct for OK
	product1_getTest, err := testQueries.GetProduct(context.Background(), product1.ProductID)
	// Check GetProduct for OK
	require.NoError(t, err)
	require.NotEmpty(t, product1_getTest)
	require.Equal(t, product1_getTest, product1)

	// Call GetProductForUpdate for OK
	product2_getTest, err := testQueries.GetProductForUpdate(context.Background(), product2.ProductID)
	// Check GetProductForUpdate for OK
	require.NoError(t, err)
	require.NotEmpty(t, product2_getTest)
	require.Equal(t, product2_getTest, product2)

	// Call UpdateProductQuantity for OK
	product1_updateTest, err := testQueries.UpdateProductQuantity(context.Background(), UpdateProductQuantityParams{
		Amount:    -5,
		ProductID: product1.ProductID,
	})
	// Check UpdateProductQuantity for OK
	require.NoError(t, err)
	require.NotEmpty(t, product1_updateTest)
	require.Equal(t, product1_getTest.Quantity-5, product1_updateTest.Quantity)
}
