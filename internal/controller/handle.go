package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-webhook-exercise/internal/domain"
)

type WxworkController struct {
	WxworkUsecase domain.IWebhook
}

func (h *WxworkController) Send(ctx *gin.Context) {
	var payload domain.WxworkPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to request body",
			"error":   err.Error(),
		})
		return
	}

	if err := h.WxworkUsecase.Send(ctx, &payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to send message",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *WxworkController) Receive(ctx *gin.Context) {
	var payload map[string]interface{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to request body",
			"error":   err.Error(),
		})
		return
	}
	if err := h.WxworkUsecase.Receive(ctx, payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to receive message",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
