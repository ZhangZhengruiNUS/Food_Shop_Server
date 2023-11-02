package handler

import (
	db "Food_Shop_Server/db/sqlc"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/* Product-count GET handle function */
func (server *Server) productCountHandler(ctx *gin.Context) {
	log.Println("================================productCountHandler: Start================================")

	var productCount int64
	var err error

	// Read frontend data
	userName := strings.TrimSpace(ctx.Query("userName"))
	log.Println("userName=", userName)

	if len(userName) == 0 {
		// If userName is empty, Get all products
		productCount, err = server.store.GetProductCount(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	} else {
		// If userName is not empty, query products of this userName
		productCount, err = server.store.GetProductCountByOwner(ctx, userName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	// Return response
	ctx.JSON(http.StatusOK, gin.H{"count": productCount})

	log.Println("================================productCountHandler: End================================")
}

/* Product-List GET handle function */
func (server *Server) productListHandler(ctx *gin.Context) {
	log.Println("================================productListHandler: Start================================")

	userName := strings.TrimSpace(ctx.Query("userName"))
	pageStr := strings.TrimSpace(ctx.Query("page"))
	pageSizeStr := strings.TrimSpace(ctx.Query("pageSize"))
	log.Println("userName=", userName)
	log.Println("page=", pageStr)
	log.Println("Demo pageSizeStr=", pageSizeStr)

	//We need both of the page and the pageSize
	if len(pageStr) == 0 || len(pageSizeStr) == 0 {
		ctx.JSON(http.StatusBadRequest, errorCustomResponse("page or pageSizeStr is empty"))
		return
	}
	pageInt, err := strconv.ParseInt(pageStr, 10, 32)
	pageSizeInt, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if len(userName) == 0 {
		// If userName is empty, Get the needed products for customer
		args := (db.GetProductListParams{
			Page:     int32(pageInt),
			Pagesize: int32(pageSizeInt),
		})

		productList, err := server.store.GetProductList(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": productList})
	} else {
		// If userName is not empty, query needed products of this userName for merchant
		args := (db.GetProductListByOwnerParams{
			ShopOwnerName: userName,
			Page:          int32(pageInt),
			Pagesize:      int32(pageSizeInt),
		})
		productList, err := server.store.GetProductListByOwner(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": productList})
	}

	log.Println("================================productListHandler: End================================")
}

/* Product GET handle function */
func (server *Server) productHandler(ctx *gin.Context) {
	log.Println("================================productHandler: Start================================")

	var err error

	// Read frontend data
	productIDStr := strings.TrimSpace(ctx.Query("productId"))
	log.Println("productIDStr=", productIDStr)

	if len(productIDStr) == 0 {
		// If productID is empty, return err
		ctx.JSON(http.StatusBadRequest, errorCustomResponse("productId is empty"))
		return
	}
	productIDInt, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// If productID is not empty, query this product
	product, err := server.store.GetProduct(ctx, productIDInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return response
	ctx.JSON(http.StatusOK, product)

	log.Println("================================productHandler: End================================")
}
