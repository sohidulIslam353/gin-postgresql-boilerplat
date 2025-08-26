package router

import (
	"gin-app/internal/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// ✅ Root route
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Gin + Bun + PostgreSQL!")
	})
	// ✅ Register API routes
	routes.RegisterRoutes(router)
	return router
}
