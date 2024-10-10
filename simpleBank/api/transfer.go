package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rpraveenkumar/Golang/db/sqlc"
)

// type transferRequest struct {
// 	FromAccountId int64  `json:"from_account_id" binding:"required,min=1"`
// 	ToAccountId   int64  `json:"to_account_id" binding:"required,min=1"`
// 	Amount        int64  `json:"Amount" binding:"required,gt=1"`
// 	Currency      string `json:"currency" binding:"required,oneof= USD EUR INR"`
// }

type transferRequest struct {
	FromAccountId int64  `json:"from_account_id" binding:"required"`
	ToAccountId   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {

	req := transferRequest{}

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	if !server.ValidAccount(ctx, req.FromAccountId, req.Currency) {
		return
	}

	if !server.ValidAccount(ctx, req.ToAccountId, req.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func (server *Server) ValidAccount(ctx *gin.Context, accountID int64, currency string) bool {

	account, err := server.store.GetAccount(ctx, accountID)

	if err != nil {
		if err == sql.ErrNoRows {

			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account %d Currency mismatch : %s vs %s ", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return false
	}

	return true
}
