package db

import (
	"Food_Shop_Server/util"
	"context"
	"time"
)

// Insert a random product in the DB
func insertRandomProduct(testQueries *Queries) (Product, error) {
	return testQueries.CreateProduct(context.Background(), CreateProductParams{
		ShopOwnerName: util.RandomString(20),
		PicPath:       util.RandomString(20),
		Describe:      util.RandomString(20),
		Price:         util.RandomFloat64(1, 1000),
		Quantity:      util.RandomInt32(1, 1000),
		ExpireTime:    time.Now().Add(24 * time.Hour),
	})
}

// Insert a random product with ownerID in the DB
func insertRandomProductWithOwner(testQueries *Queries, shopOwnerName string) (Product, error) {
	return testQueries.CreateProduct(context.Background(), CreateProductParams{
		ShopOwnerName: shopOwnerName,
		PicPath:       util.RandomString(20),
		Describe:      util.RandomString(20),
		Price:         util.RandomFloat64(1, 1000),
		Quantity:      util.RandomInt32(1, 1000),
		ExpireTime:    time.Now().Add(24 * time.Hour),
	})
}

// Delete a product in the DB
func deleteTestProduct(testQueries *Queries, productID int64) error {
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
