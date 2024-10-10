package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rpraveenkumar/Golang/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", ValidCurrency)
	}

	server := &Server{store: store}
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)
	router.POST("/transfers/", server.createTransfer)

	server.router = router
	return server

}
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}

}

// start the server with speicfu aaddress
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
