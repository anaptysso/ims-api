package routes

import (
	echo "github.com/labstack/echo/v4"
)

// Base works for all the routes
type Base struct {
	R *echo.Group
}
