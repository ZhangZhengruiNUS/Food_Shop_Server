package handler

import (
	db "Food_Shop_Server/db/sqlc"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/* Product-add received data */
type productAddRequest struct {
	ShopOwnerName string  `json:"userName"`
	PicPath       string  `json:"picPath"`
	Describe      string  `json:"describe"`
	Price         float64 `json:"price"`
	Quantity      int32   `json:"quantity"`
	ExpireTime    string  `json:"expireTime"`
}

/* Product-add Post handle function */
func (server *Server) productAddHandler(ctx *gin.Context) {
	log.Println("================================productAddHandler: Start================================")

	var req productAddRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("ShopOwnerName=", req.ShopOwnerName)
	log.Println("PicPath=", req.PicPath)
	log.Println("Describe=", req.Describe)
	log.Println("Price=", req.Price)
	log.Println("Quantity=", req.Quantity)
	log.Println("ExpireTime=", req.ExpireTime)
	if len(req.ShopOwnerName) == 0 || len(req.Describe) == 0 || req.Price == 0 || req.Quantity == 0 || len(req.ExpireTime) == 0 {
		log.Println("Incomplete data")
		ctx.JSON(http.StatusBadRequest, errorCustomResponse("Incomplete data"))
		return
	}
	expireTime, err := time.Parse("20060102", req.ExpireTime)
	if err != nil {
		log.Println("Convert time wrong")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Start database transaction
	err = server.store.ExecTx(ctx, func(q *db.Queries) error {
		// CreateProduct
		product, err := server.store.CreateProduct(ctx, db.CreateProductParams{
			ShopOwnerName: req.ShopOwnerName,
			PicPath:       req.PicPath,
			Describe:      req.Describe,
			Price:         req.Price,
			Quantity:      req.Quantity,
			ExpireTime:    expireTime,
		})
		if err != nil {
			return err
		}
		log.Println("Product Created")

		ctx.JSON(http.StatusOK, product)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("================================productAddHandler: End================================")
}

/* Product-delete Delete handle function */
func (server *Server) productDeleteHandler(ctx *gin.Context) {
	log.Println("================================productDeleteHandler: Start================================")

	// Read frontend data
	productIDStr := strings.TrimSpace(ctx.Param("id"))
	if len(productIDStr) == 0 {
		ctx.JSON(http.StatusBadRequest, errorCustomResponse("productIDStr is empty"))
		return
	}
	productIDInt, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	log.Println("ProductID=", productIDInt)

	// Start database transaction
	err = server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Check product exits
		_, err := server.store.GetProductForUpdate(ctx, productIDInt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If product not exists, return err
				log.Println("Product not exists")
				ctx.JSON(http.StatusNotFound, errorCustomResponse("ProductID:["+strconv.FormatInt(productIDInt, 10)+"]not exists!"))
				return nil
			} else {
				return err
			}
		}
		// If product exists, delete it
		err = server.store.DeleteProduct(ctx, productIDInt)
		if err != nil {
			return err
		}
		log.Println("ProductID:[" + strconv.FormatInt(productIDInt, 10) + "] Deleted")

		ctx.JSON(http.StatusOK, nil)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("================================productDeleteHandler: End================================")
}

/* Product-buy received data */
type productBuyRequest struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

/* Product-buy Post handle function */
func (server *Server) productBuyHandler(ctx *gin.Context) {
	log.Println("================================productBuyHandler: Start================================")

	// Read frontend data
	var req productBuyRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	log.Println("ProductID=", req.ProductID)
	log.Println("Quantity=", req.Quantity)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Check product exits
		_, err := server.store.GetProductForUpdate(ctx, req.ProductID)
		if err != nil {
			if err == sql.ErrNoRows {
				// If product not exists, return err
				log.Println("Product not exists")
				ctx.JSON(http.StatusNotFound, errorCustomResponse("ProductID:["+strconv.FormatInt(req.ProductID, 10)+"]not exists!"))
				return nil
			} else {
				return err
			}
		}
		// If product exists, update quantity
		product, err := server.store.UpdateProductQuantity(ctx, db.UpdateProductQuantityParams{
			Amount:    -req.Quantity,
			ProductID: req.ProductID,
		})
		if err != nil {
			return err
		}
		log.Println("ProductID:[" + strconv.FormatInt(product.ProductID, 10) + "] Quantity:[" + strconv.FormatInt(int64(product.Quantity), 10) + "] Updated")

		ctx.JSON(http.StatusOK, product)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	log.Println("================================productBuyHandler: End================================")
}
