package bootstrap

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-webhook-exercise/internal/controller"
	"github.com/rs/zerolog"
)

type BootStrap struct {
	engin *gin.Engine
}

func NewBootStrap() *BootStrap {
	return &BootStrap{
		engin: gin.Default(),
	}
}

func (b *BootStrap) Run() error {
	_ = b.engin.SetTrustedProxies(nil)

	if err := b.MapRoutes(); err != nil {
		return err
	}

	if err := b.engin.Run(":8899"); err != nil {
		return err
	}

	return nil
}

func (b *BootStrap) MapRoutes() error {
	v1 := b.engin.Group("/api/v1")
	log := zerolog.New(os.Stdout)
	controller.Routes(v1, &log)
	return nil
}
