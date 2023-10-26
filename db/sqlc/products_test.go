package db

import (
	"Food_Shop_Server/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {
	// Create some test data in the DB
	product1, err := insertRandomProduct()
	require.NoError(t, err)
	product2, err := insertRandomProduct()
	require.NoError(t, err)
	product3, err := insertRandomProductWithID(product1.ShopOwnerID)
	require.NoError(t, err)

	// Call GetProductCount
	count, err := testQueries.GetProductCount(context.Background())
	// Check GetProductCount
	require.NoError(t, err)
	require.NotEmpty(t, count)
	require.Equal(t, int64(3), count)

	// Call GetProductCountByOwner
	count, err = testQueries.GetProductCountByOwner(context.Background(), product1.ShopOwnerID)
	// Check GetProductCountByOwner
	require.NoError(t, err)
	require.NotEmpty(t, count)
	require.Equal(t, int64(2), count)

	// Delete test data
	deleteRandomProduct(product1.ProductID)
	deleteRandomProduct(product2.ProductID)
	deleteRandomProduct(product3.ProductID)
}

// Insert a random product in the DB
func insertRandomProduct() (Product, error) {
	return testQueries.CreateProduct(context.Background(), CreateProductParams{
		ShopOwnerID: util.RandomInt64(1, 1000),
		PicPath:     util.RandomString(20),
		Describe:    util.RandomString(20),
		Price:       util.RandomInt32(1, 1000),
		Quantity:    util.RandomInt32(1, 1000),
	})
}

// Insert a random product with ownerID in the DB
func insertRandomProductWithID(shopOwnerID int64) (Product, error) {
	return testQueries.CreateProduct(context.Background(), CreateProductParams{
		ShopOwnerID: shopOwnerID,
		PicPath:     util.RandomString(20),
		Describe:    util.RandomString(20),
		Price:       util.RandomInt32(1, 1000),
		Quantity:    util.RandomInt32(1, 1000),
	})
}

// Delete a product in the DB
func deleteRandomProduct(productID int64) error {
	return testQueries.DeleteProduct(context.Background(), productID)
}
