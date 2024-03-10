package handlers

import (
	"fmt"
	"net/http"

	"github.com/harish876/go-templ-htmx-exploration/models"
	"github.com/harish876/go-templ-htmx-exploration/services"
	"github.com/harish876/go-templ-htmx-exploration/views/grid"
	"github.com/labstack/echo/v4"
)

func GridHandler(c echo.Context) error {
	return Render(c, http.StatusOK, grid.Grid(services.Data))
}

func GridRowHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(services.Data, id)
	return Render(c, http.StatusOK, grid.RenderRow(row))
}

func UpdateGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	var newData = models.GridDataRow{
		Id:       id,
		Status:   c.FormValue("status"),
		Position: c.FormValue("position"),
	}
	updatedRow := services.UpdateById(services.Data, id, newData)
	fmt.Println(updatedRow)
	return Render(c, http.StatusOK, grid.RenderRow(updatedRow))
}

func DeleteGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	services.DeleteById(id)
	return c.NoContent(http.StatusOK)
}

func RenderEditGridHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(services.Data, id)
	return Render(c, http.StatusOK, grid.EditRow(row))
}
