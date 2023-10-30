package handler

import (
	db "Food_Shop_Server/db/sqlc"
	"Food_Shop_Server/util"
	"bytes"
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

/* Check the expect and actual HTTP body */
func requireBodyMatch(t *testing.T, body *bytes.Buffer, expect gin.H) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var actual gin.H

	err = json.Unmarshal(data, &actual)
	require.NoError(t, err)
	require.Equal(t, expect, actual)
}

// Generate random GetProductListRows
func creatRandomGetProductListRow(count int) []db.GetProductListRow {
	var getProductListRows []db.GetProductListRow
	for i := 0; i < count; i++ {
		getProductListRows = append(getProductListRows, db.GetProductListRow{
			ProductID: util.RandomInt64(1, 100),
			Describe:  util.RandomString(20),
			PicPath:   util.RandomString(20),
		})
	}
	return getProductListRows
}

// Generate random GetProductListByOwnerRows
func creatRandomGetProductListByOwnerRow(count int) []db.GetProductListByOwnerRow {
	var getProductListByOwnerRows []db.GetProductListByOwnerRow
	for i := 0; i < count; i++ {
		getProductListByOwnerRows = append(getProductListByOwnerRows, db.GetProductListByOwnerRow{
			ProductID: util.RandomInt64(1, 100),
			Describe:  util.RandomString(20),
			PicPath:   util.RandomString(20),
		})
	}
	return getProductListByOwnerRows
}

// Generate random products
func creatRandomProducts(count int) []db.Product {
	var products []db.Product
	for i := 0; i < count; i++ {
		products = append(products, db.Product{
			ProductID:     util.RandomInt64(1, 100),
			ShopOwnerName: util.RandomString(20),
			PicPath:       util.RandomString(20),
			Describe:      util.RandomString(20),
			Price:         util.RandomFloat64(1, 1000),
			Quantity:      util.RandomInt32(10, 1000),
			ExpireTime:    time.Now().Add(24 * time.Hour),
			CreateTime:    time.Now(),
		})
	}
	return products
}

// Convert struct to gin.H
func structToGinH(s interface{}) gin.H {
	bytes, _ := json.Marshal(s)

	var h gin.H
	_ = json.Unmarshal(bytes, &h)

	return h
}

// Convert map to io.Reader
func ginHToIoReader(h gin.H) io.Reader {
	jsonData, _ := json.Marshal(h)

	reader := bytes.NewBuffer(jsonData)
	return reader
}

// Generate random CreateProductParams
func createRandomProductAddRequest() productAddRequest {
	return productAddRequest{
		ShopOwnerName: util.RandomString(20),
		PicPath:       util.RandomString(20),
		Describe:      util.RandomString(20),
		Price:         util.RandomFloat64(1, 1000),
		Quantity:      util.RandomInt32(10, 1000),
		ExpireTime:    time.Now().Add(24 * time.Hour).Format("20060102"),
	}
}

// Generate product by CreateProductParams
func createProductByRequest(req productAddRequest) db.Product {
	expireTime, _ := time.Parse("20060102", req.ExpireTime)
	return db.Product{
		ProductID:     util.RandomInt64(1, 100),
		ShopOwnerName: req.ShopOwnerName,
		PicPath:       req.PicPath,
		Describe:      req.Describe,
		Price:         req.Price,
		Quantity:      req.Quantity,
		ExpireTime:    expireTime,
		CreateTime:    time.Now(),
	}
}
