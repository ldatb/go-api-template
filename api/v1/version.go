package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Version represents the API version information.
type Version struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Date   string `json:"date"`
}

// GetVersion handles the request to get the API version information.
func GetVersion(c echo.Context) error {
	version := Version{
		ID:     "v0.1.0",
		Status: "experimental",
		Date:   "2024-07-25",
	}
	return c.JSON(http.StatusOK, map[string]Version{"version": version})
}
