package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/a-h/templ"
	"github.com/harish876/go-templ-htmx-exploration/db"
	"github.com/harish876/go-templ-htmx-exploration/models"
	"github.com/harish876/go-templ-htmx-exploration/views/components"
	"github.com/harish876/go-templ-htmx-exploration/views/form"
	"github.com/harish876/go-templ-htmx-exploration/views/grid"
	"github.com/harish876/go-templ-htmx-exploration/views/home"
	"github.com/labstack/echo/v4"
)

/* TODO: MAKE HANDLERS AND STRUCTURE */

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func FormSubmissionHandler(c echo.Context) error {
	formValues := form.FormValues{
		FirstName:            c.FormValue("first_name"),
		LastName:             c.FormValue("last_name"),
		Email:                c.FormValue("email"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("password_confirmation"),
		MarketingAccept:      c.FormValue("marketing_accept"),
	}

	var formErrors form.FormErrors
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailRegex)

	fmt.Println(formValues)
	if formValues.Password != formValues.PasswordConfirmation {
		formErrors.Password = "Password and Password Confirmation must match"
	}
	if ok := regex.MatchString(formValues.Email); !ok {
		formErrors.Email = "Invalid email format"
	}

	if formErrors.Email != "" || formErrors.Password != "" {
		fmt.Println(formErrors)
		return Render(c, 422, form.Form(formValues, formErrors, false, ""))
	}

	return Render(c, http.StatusOK, form.Form(form.FormValues{}, form.FormErrors{}, true, "Success"))
}

func FormHandler(c echo.Context) error {
	return Render(c, http.StatusOK, form.Form(form.FormValues{}, form.FormErrors{}, false, ""))
}

func GridHandler(c echo.Context) error {
	return Render(c, http.StatusOK, grid.Grid(db.Data))
}

func GridRowHandler(c echo.Context) error {
	id := c.Param("id")
	row := db.FilterById(db.Data, id)
	return Render(c, http.StatusOK, grid.RenderRow(row))
}

func UpdateGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	var newData = models.GridDataRow{
		Id:       id,
		Status:   c.FormValue("status"),
		Position: c.FormValue("position"),
	}
	updatedRow := db.UpdateById(db.Data, id, newData)
	fmt.Println(updatedRow)
	return Render(c, http.StatusOK, grid.RenderRow(updatedRow))
}

func DeleteGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	db.DeleteById(id)
	return c.NoContent(http.StatusOK)
}

func RenderEditGridHandler(c echo.Context) error {
	id := c.Param("id")
	row := db.FilterById(db.Data, id)
	return Render(c, http.StatusOK, grid.EditRow(row))
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, home.Home())
}

func main() {
	app := echo.New()
	app.Static("/", "assets")

	// Set custom HTTP error handler for 404 Not Found
	app.HTTPErrorHandler = func(err error, c echo.Context) {
		// Check if the error is a 404 Not Found error
		if h, ok := err.(*echo.HTTPError); ok {
			if h.Code == http.StatusNotFound {
				Render(c, http.StatusNotFound, components.NotFound())
			}
		}

		// For other errors, fall back to Echo's default error handler
		app.DefaultHTTPErrorHandler(err, c)
	}

	app.GET("/", HomeHandler)
	app.GET("/form", FormHandler)
	app.POST("/form", FormSubmissionHandler)
	app.GET("/grid", GridHandler)

	app.GET("/grid/:id", GridRowHandler)
	app.PUT("/grid/:id", UpdateGridRowHandler)
	app.DELETE("/grid/:id", DeleteGridRowHandler)

	app.GET("/grid/edit/:id", RenderEditGridHandler)

	app.Logger.Fatal(app.Start(":42069"))
}
