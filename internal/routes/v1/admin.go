package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(rg *gin.RouterGroup) {
	// Admin routes go here
	rg.GET("/dashboard", func(c *gin.Context) {
		c.String(200, "Welcome to the Admin Dashboard")
	})
}
