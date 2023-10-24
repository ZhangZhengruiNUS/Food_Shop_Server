package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Login received data */
type loginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

/* Login handle function */
func (server *Server) loginHandler(ctx *gin.Context) {
	fmt.Println("================================loginHandler: Start================================")

	var req loginRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("username=", req.UserName)

	// Get data from database
	user, err := server.store.GetUserByName(ctx, req.UserName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Verify password and return response
	if user.Password == req.Password {
		fmt.Println("password right")
		ctx.JSON(http.StatusOK, user)
	} else {
		fmt.Println("password wrong")
		ctx.JSON(http.StatusUnauthorized, commonResponse("userName & password denied"))
	}

	fmt.Println("================================loginHandler: End================================")
}
