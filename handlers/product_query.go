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
	ctx.JSON(http.StatusOK, productCount)

	log.Println("================================productCountHandler: End================================")
}

// merchant get product list
func (server *Server) productListHandler(ctx *gin.Context) {
	log.Println("================================merchantProductList: Start================================")

	userIDStr := strings.TrimSpace(ctx.Query("userId"))
	pageStr := strings.TrimSpace(ctx.Query("page"))
	pageSizeStr := strings.TrimSpace(ctx.Query("pageSize"))
	log.Println("userID=", userIDStr)
	log.Println("page=", pageStr)
	log.Println("pageSizeStr=", pageSizeStr)

	//We need both of the page and the pageSize
	if len(pageStr) == 0 || len(pageSizeStr) == 0 {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	pageInt, err := strconv.ParseInt(pageStr, 10, 32)
	pageSizeInt, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if len(userIDStr) == 0 {
		// If userId is empty, Get the needed products for customer
		productList, err = server.store.GetProductList(ctx, pageInt, pageSizeInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		// If userId is not empty, query needed products of this userId for merchant
		userIDInt, err := strconv.ParseInt(userIDStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		productList, err = server.store.GetProductListByOwner(ctx, userIDInt, pageInt, pageSizeInt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	// Return response
	ctx.JSON(http.StatusOK, productList)

	log.Println("================================merchantProductList: End================================")
}
