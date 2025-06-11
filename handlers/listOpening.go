package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}