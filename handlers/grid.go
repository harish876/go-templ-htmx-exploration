package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/harish876/go-templ-htmx-exploration/models"
	"github.com/harish876/go-templ-htmx-exploration/services"
	"github.com/harish876/go-templ-htmx-exploration/views/grid"
	"github.com/labstack/echo/v4"
)

var (
	DEFAULT_PAGE_SIZE = 1
	DEFAULT_LIMIT     = 10
)

func GridHandler(c echo.Context) error {
	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	limit, limitErr := strconv.Atoi(c.QueryParam("limit"))

	if pageErr != nil {
		page = DEFAULT_PAGE_SIZE
	}

	if limitErr != nil {
		limit = DEFAULT_LIMIT
	}

	totalPages := len(services.Data)/limit + 1
	offset := (page - 1) * limit

	var paginatedData []models.GridDataRow
	if offset < len(services.Data) {
		end := offset + limit
		if end > len(services.Data) {
			end = len(services.Data)
		}
		paginatedData = services.Data[offset:end]
	} else {
		paginatedData = models.GridData{}
	}

	return Render(c, http.StatusOK, grid.Grid(paginatedData, totalPages))
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
