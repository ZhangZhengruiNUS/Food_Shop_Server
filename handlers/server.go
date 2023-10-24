package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

// server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	router.POST("/login", server.loginHandler)

	router.POST("/signup", server.signupHandler)

	router.GET("/catalog", server.catalogHandler)
	router.GET("/catalog/user", server.catalogUserHandler)
	router.POST("/catalog/buy", server.catalogBuyHandler)

	router.GET("/game/diff", server.gameDiffHandler)
	router.GET("/game/equip", server.gameEquipQueryHandler)
	router.POST("/game/equip", server.gameEquipUpdateHandler)
	router.POST("/game/end", server.gameEndHandler)

	server.router = router
	return server
}

// start runs the HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// handle error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// handle common response
func commonResponse(msg string) gin.H {
	return gin.H{"error": msg}
}
