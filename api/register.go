package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jpmoraess/appointment-api/db/sqlc"
	"net/http"
)

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.RegisterTxParams{
		Name:     req.Name,
		Password: req.Password,
	}

	result, err := server.store.RegisterTx(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, result)
}
