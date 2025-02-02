package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getTenantRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetTenant(ctx *gin.Context) {
	var req getTenantRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tenant, err := server.store.GetTenant(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tenant)
}
