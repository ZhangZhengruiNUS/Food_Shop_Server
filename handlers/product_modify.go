package handler

import (
	db "Food_Shop_Server/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/* Product-add received data */
type productAddRequest struct {
	ShopOwnerID int64     `json:"shopOwnerId"`
	PicPath     string    `json:"picPath"`
	Describe    string    `json:"describe"`
	Price       int32     `json:"price"`
	Quantity    int32     `json:"quantity"`
	CreateTime  time.Time `json:"createTime"`
}

type productDeleteRequest struct {
	ShopOwnerID int64 `json:"shop_owner_id"`
	Product_ID  int64 `json:"product_ID"`
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
	fmt.Println("Shop_owner_id=", req.ShopOwnerID)
	fmt.Println("Pic_path=", req.PicPath)
	fmt.Println("ProductInfo=", req.Describe)
	fmt.Println("Price=", req.Price)
	fmt.Println("quantity=", req.Quantity)
	fmt.Println("ExpireTime=", req.CreateTime)
	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {

		//todo test cases in db.go
		//user, err := server.store.GetUserById(ctx, req.Shop_owner_id)

		// Read item's price
		//item, err := server.store.GetItem(ctx, req.ItemID)

		// Update user's credit

		// Read quantity

		// If ID not exists, create it
		server.store.CreateItem(ctx, db.CreateItemParams{
			ShopOwnerID: req.ShopOwnerID,
			// ItemID:   req.ItemID,
			Quantity:   req.Quantity,
			Price:      req.Price,
			CreateTime: req.CreateTime,
			PicPath:    req.PicPath,
			Describe:   req.Describe,
		})
		// if err != nil {
		// 	return err
		// }
		fmt.Println("add Created")

		ctx.JSON(http.StatusOK, req.ShopOwnerID)
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

	var reqDelete productDeleteRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&reqDelete); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("Product_ID=", reqDelete.Product_ID)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {

		//todo test cases in db.go
		//user, err := server.store.GetUserById(ctx, req.Shop_owner_id)

		// Read item's price
		item, err := server.store.GetItem(ctx, reqDelete.Product_ID)

		// Update user's credit

		// Access control
		if reqDelete.ShopOwnerID == item.ShopOwnerID {

			// If ID not exists, create it
			server.store.DeleteItem(ctx, reqDelete.Product_ID)
			if err != nil {
				return err
			}
			fmt.Println("Product Delete")

		}
		ctx.JSON(http.StatusOK, reqDelete.ShopOwnerID)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================productDeleteHandler: End================================")
}
