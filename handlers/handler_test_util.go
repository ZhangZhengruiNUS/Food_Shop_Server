package handler

import (
	db "Food_Shop_Server/db/sqlc"
	"Food_Shop_Server/util"
	"bytes"
	"encoding/json"
	"io"
	"testing"

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
