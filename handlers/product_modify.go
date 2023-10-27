// Leilei
package handler
/* Catalog-buy received data */
type catalogBuyRequest struct {
	UserID int64 `json:"userId"`
	ItemID int64 `json:"itemId"`
  ProductName 
}


/* Catalog-add Post handle function */
func (server *Server) catalogBuyHandler(ctx *gin.Context) {
	fmt.Println("================================catalogBuyHandler: Start================================")

	var req productAddRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("UserID=", req.UserID)
	fmt.Println("ItemID=", req.ItemID)
