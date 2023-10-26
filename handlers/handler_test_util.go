package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

/* Check the expect and actual HTTP body */
func requireBodyMatch(t *testing.T, body *bytes.Buffer, result gin.H) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotCount gin.H

	err = json.Unmarshal(data, &gotCount)
	require.NoError(t, err)
	require.Equal(t, result, gotCount)
}
