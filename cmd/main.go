package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/harish876/go-templ-htmx-exploration/views/form"
	"github.com/harish876/go-templ-htmx-exploration/views/home"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func FormHandler(c echo.Context) error {
	return Render(c, http.StatusOK, form.Form())
}

func Homeandler(c echo.Context) error {
	return Render(c, http.StatusOK, home.Home())
}

func main() {
	app := echo.New()
	app.Static("/", "assets")
	app.GET("/", Homeandler)
	app.GET("/form", FormHandler)
	app.Logger.Fatal(app.Start(":42069"))
}
