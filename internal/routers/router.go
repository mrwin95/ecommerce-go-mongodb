package routers

import (
	"mrwin95/go-ecommerce-backend-api/internal/controller"
	"mrwin95/go-ecommerce-backend-api/internal/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	// r := gin.Default() // Logger and Recovery are enabled by default
	r := gin.New()

	// Security header
	expectedHost := "localhost:8080"

	// Setup Security Headers
	r.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	//om router
	r.Use(middlewares.Logger())

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", controller.Pong)

		v2.GET("/user/:id", controller.GetUserById)

		v2.POST("/signin", controller.SignIn)
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
	return r
}
