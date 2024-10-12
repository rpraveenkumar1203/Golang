package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/rpraveenkumar/Golang/db/sqlc"
	"github.com/rpraveenkumar/Golang/db/utils"

	"github.com/rpraveenkumar/Golang/token"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.Token_symmetric_key)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", ValidCurrency)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)
	router.POST("/transfers/", server.createTransfer)

	server.router = router
	return server, nil

}
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}

}

// start the server with speicfu aaddress
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
