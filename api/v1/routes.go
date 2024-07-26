package v1

import (
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers the version 1 API routes.
func RegisterRoutes(router *echo.Echo) {
	// Define v1 group
	v1 := router.Group("/api/v1")

	// Define routes
	v1.GET("", GetVersion)
}
