package routes

import (
	"github.com/harish876/go-templ-htmx-exploration/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterFormRoutes(e *echo.Echo) {
	// Route grouping for form
	apiGroup := e.Group("/form")

	apiGroup.GET("", handlers.FormHandler)
	apiGroup.POST("", handlers.FormSubmissionHandler)
}
