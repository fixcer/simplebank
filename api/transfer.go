package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/fixcer/simplebank/api/model"
	db "github.com/fixcer/simplebank/db/sqlc"
	"github.com/fixcer/simplebank/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createTransfer(ctx *gin.Context) {
	var req model.TransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	fromAccount, ok := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !ok {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("from_account does not belong to the authenticated user")))
		return
	}

	_, ok = server.validAccount(ctx, req.ToAccountID, req.Currency)
	if !ok {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)))
		return account, false
	}

	return account, true
}
