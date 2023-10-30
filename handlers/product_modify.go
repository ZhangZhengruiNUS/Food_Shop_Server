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
	ShopOwnerName string    `json:"userName"`
	PicPath       string    `json:"picPath"`
	Describe      string    `json:"describe"`
	Price         float64   `json:"price"`
	Quantity      int32     `json:"quantity"`
	ExpireTime    time.Time `json:"expireTime"`
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
	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// CreateProduct
		product, err := server.store.CreateProduct(ctx, db.CreateProductParams{
			ShopOwnerName: req.ShopOwnerName,
			PicPath:       req.PicPath,
			Describe:      req.Describe,
			Price:         req.Price,
			Quantity:      req.Quantity,
			ExpireTime:    req.ExpireTime,
		})
		if err != nil {
			return err
		}
		log.Println("add Created")

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

/* Product-delete delete handle function */
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
		_, err := server.store.GetProduct(ctx, productIDInt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If product not exists, return err
				log.Println("User not exists")
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
