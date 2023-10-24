package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Signup received data */
type signupRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

/* Signup handle function */
func (server *Server) signupHandler(ctx *gin.Context) {
	fmt.Println("================================signupHandler: Start================================")

	var req signupRequest

	// Read frontend data
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("username=", req.UserName)

	// Start database transaction
	err := server.store.ExecTx(ctx, func(q *db.Queries) error {
		// Read user
		user, err := server.store.GetUserByName(ctx, req.UserName)
		if err != nil {
			if err == sql.ErrNoRows {
				// If username not exists, create it
				user, err = server.store.CreateUser(ctx, db.CreateUserParams{
					UserName: req.UserName,
					Password: req.Password,
					RoleType: 0,
				})
				if err != nil {
					return err
				}
				fmt.Println("User Created")
			} else {
				return err
			}
		} else {
			// If username exists, return err
			ctx.JSON(http.StatusConflict, commonResponse("UserName:["+user.UserName+"]has already been signed up!"))
			return nil
		}

		ctx.JSON(http.StatusOK, user)
		return nil
	})

	// Return response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("================================signupHandler: End================================")
}
