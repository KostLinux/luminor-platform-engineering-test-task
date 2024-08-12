package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "OK",
		"message": message,
	})
}

func BadRequestError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    "BAD_REQUEST",
		"message": "Invalid request. Please check your JSON fields and try again.",
		"error":   err.Error(),
	})
}

func ValidationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    "BAD_REQUEST",
		"message": "Failed to process payload.",
		"error":   err.Error(),
	})
}

func InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    "INTERNAL_SERVER_ERROR",
		"message": "Failed to process payload.",
		"error":   err.Error(),
	})
}
