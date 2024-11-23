package routers

import (
	"mrwin95/go-ecommerce-backend-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	r := gin.Default()
	v2 := r.Group("/v2")
	{
		v2.GET("/ping", controller.Pong)
	}

	return r
}
