package handler

import (
	db "Food_Shop_Server/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Product-add received data */
type productAddRequest struct {
	Shop_owner_id int64   `json:"Shop_owner_id"`
	Pic_path      string  `json:"pic_path"`
	ProductInfo   string  `json:"productInfo"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	ExpireTime    string  `json:"expireTime"`
	Product_ID    int64
}

/* Product-add Post handle function */
func (server *Server) productAddHandler(ctx *gin.Context) {
	fmt.Println("================================productAddHandler: Start================================")

	var req productAddRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("Shop_owner_id=", req.Shop_owner_id)
	fmt.Println("Pic_path=", req.Pic_path)
	fmt.Println("ProductInfo=", req.ProductInfo)
	fmt.Println("Price=", req.Price)
	fmt.Println("quantity=", req.Quantity)
	fmt.Println("ExpireTime=", req.ExpireTime)
	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {

		//todo test cases in db.go
		user, err := server.store.GetUserById(ctx, req.Shop_owner_id)

		// Read item's price
		// item, err := server.store.GetItem(ctx, req.ItemID)

		// Update user's credit

		// Read quantity
		if err != nil {
			if err == sql.ErrNoRows {
				// If ID not exists, create it
				server.store.Createproduct(ctx, db.CreateproductParams{
					Shop_owner_id: user.Shop_owner_id,
					// ItemID:   req.ItemID,
					Quantity:    req.Quantity,
					Price:       req.Price,
					ExpireTime:  req.ExpireTime,
					Pic_path:    req.Pic_path,
					ProductInfo: req.ProductInfo,
				})
				if err != nil {
					return err
				}
				fmt.Println("quantity Created")
			} else {
				return err
			}
		}
		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================productAddHandler: End================================")
}

/* Product-delete delete handle function */
func (server *Server) productDeleteHandler(ctx *gin.Context) {
	fmt.Println("================================productDeleteHandler: Start================================")

	var req productAddRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("Product_ID=", req.Product_ID)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {

		//todo test cases in db.go
		user, err := server.store.GetUserById(ctx, req.Shop_owner_id)

		// Read item's price
		item, err := server.store.GetItem(ctx, req.ItemID)

		// Update user's credit

		// Read quantity
		if err != nil {
			if err == sql.ErrNoRows {
				// If ID not exists, create it
				server.store.Deleteproduct(ctx, item)
				if err != nil {
					return err
				}
				fmt.Println("Product Delete")
			} else {
				return err
			}
		}
		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================productDeleteHandler: End================================")
}
