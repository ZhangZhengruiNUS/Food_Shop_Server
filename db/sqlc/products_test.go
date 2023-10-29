package db

import (
	"Food_Shop_Server/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct(t *testing.T) {

	var product1 Product
	var product2 Product
	var product3_1 Product
	var product4_1 Product
	var product5_1 Product
	var product6_2 Product
	var product7_1 Product
	var product8 Product
	var err error

	// Delete test data (when test is over)
	defer func() {
		deleteTestProduct(product1.ProductID)
		deleteTestProduct(product2.ProductID)
		deleteTestProduct(product3_1.ProductID)
		deleteTestProduct(product4_1.ProductID)
		deleteTestProduct(product5_1.ProductID)
		deleteTestProduct(product6_2.ProductID)
		deleteTestProduct(product7_1.ProductID)
		deleteTestProduct(product8.ProductID)
	}()

	// Create some test data in the DB
	product1, err = insertRandomProduct()
	require.NoError(t, err)
	product2, err = insertRandomProduct()
	require.NoError(t, err)
	product3_1, err = insertRandomProductWithOwner(product1.ShopOwnerName)
	require.NoError(t, err)
	product4_1, err = insertRandomProductWithOwner(product1.ShopOwnerName)
	require.NoError(t, err)
	product5_1, err = insertRandomProductWithOwner(product1.ShopOwnerName)
	require.NoError(t, err)
	product6_2, err = insertRandomProductWithOwner(product2.ShopOwnerName)
	require.NoError(t, err)
	product7_1, err = insertRandomProductWithOwner(product1.ShopOwnerName)
	require.NoError(t, err)
	product8, err = insertRandomProduct()
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
}

// Insert a random product in the DB
func insertRandomProduct() (Product, error) {
	return testQueries.CreateProduct(context.Background(), CreateProductParams{
		ShopOwnerName: util.RandomString(20),
		PicPath:       util.RandomString(20),
		Describe:      util.RandomString(20),
		Price:         util.RandomInt32(1, 1000),
		Quantity:      util.RandomInt32(1, 1000),
	})
}

// Insert a random product with ownerID in the DB
func insertRandomProductWithOwner(shopOwnerName string) (Product, error) {
	return testQueries.CreateProduct(context.Background(), CreateProductParams{
		ShopOwnerName: shopOwnerName,
		PicPath:       util.RandomString(20),
		Describe:      util.RandomString(20),
		Price:         util.RandomInt32(1, 1000),
		Quantity:      util.RandomInt32(1, 1000),
	})
}

// Delete a product in the DB
func deleteTestProduct(productID int64) error {
	if productID == 0 {
		return nil
	}

	return testQueries.DeleteProduct(context.Background(), productID)
}

// Create a var of struct createProductListRow
func createProductListRow(product Product) GetProductListRow {
	return GetProductListRow{
		ProductID: product.ProductID,
		Describe:  product.Describe,
		PicPath:   product.PicPath}
}

// Create a var of struct createProductListRow
func createProductListByOwnerRow(product Product) GetProductListByOwnerRow {
	return GetProductListByOwnerRow{
		ProductID: product.ProductID,
		Describe:  product.Describe,
		PicPath:   product.PicPath}
}
