package main

import (
	"github.com/harish876/go-templ-htmx-exploration/handlers"
	"github.com/harish876/go-templ-htmx-exploration/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")

	// Set custom HTTP error handler for 404 Not Found
	app.HTTPErrorHandler = handlers.NotFoundErrorHandler

	routes.RegisterHomeRoutes(app)
	routes.RegisterGridRoutes(app)
	routes.RegisterFormRoutes(app)

	app.Logger.Fatal(app.Start(":42070"))
}
