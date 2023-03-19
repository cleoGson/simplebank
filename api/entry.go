package api

import (
	"database/sql"
	"net/http"

	db "github.com/cleoGson/simplebank/db"
	"github.com/gin-gonic/gin"
)

type CreateEntryRequest struct {
	AccountID int64 `json:"account_id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`
}

func (server *Server) createEntry(ctx *gin.Context) {

	var req CreateEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	Entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Entry)
}

type getEntryRequest struct {
	ID int64 `uri:"id" binding: "required, min=1 "`
}

func (server *Server) getEntry(ctx *gin.Context) {

	var req getEntryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	Entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Entry)
}

type ListEntryRequest struct {
	PageID   int32 `form:"page_id" binding: "required, min=1 "`
	pageSize int32 `form:"page_size" binding: "required, min=5, max=100"`
}

func (server *Server) listEntries(ctx *gin.Context) {

	var req ListEntryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListEntriesParams{
		Limit:  req.pageSize,
		Offset: (req.PageID - 1) * req.pageSize,
	}

	Entries, err := server.store.ListEntries(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Entries)
}
