package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-webhook-exercise/internal/usecase"
	"github.com/rs/zerolog"
)

func Routes(r *gin.RouterGroup, logger *zerolog.Logger) {
	usecase := usecase.NewWxworkUsecase(logger)
	handle := WxworkController{
		WxworkUsecase: usecase,
	}

	route := r.Group("/wxwork")
	{
		route.POST("", handle.Send)
		route.POST("/receive", handle.Receive)
	}
}
