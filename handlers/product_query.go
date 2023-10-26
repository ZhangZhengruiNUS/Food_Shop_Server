package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/* Product-count handle function */
func (server *Server) productCountHandler(ctx *gin.Context) {
	log.Println("================================productCountHandler: Start================================")

	var productCount int64
	var err error

	// Read frontend data
	userIDStr := strings.TrimSpace(ctx.Query("userId"))
	log.Println("userID=", userIDStr)

	if len(userIDStr) == 0 {
		// If userId is empty, Get all products
		productCount, err = server.store.GetProductCount(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		// If userId is not empty, query products of this userId
		userIDInt, err := strconv.ParseInt(userIDStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		productCount, err = server.store.GetProductCountByOwner(ctx, userIDInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	// Return response
	ctx.JSON(http.StatusOK, gin.H{"count": productCount})

	log.Println("================================productCountHandler: End================================")
}
