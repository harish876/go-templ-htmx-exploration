package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/harish876/go-templ-htmx-exploration/models"
	"github.com/harish876/go-templ-htmx-exploration/services"
	"github.com/harish876/go-templ-htmx-exploration/views/grid"
	"github.com/labstack/echo/v4"
)

var (
	DEFAULT_PAGE_SIZE = 1
	DEFAULT_LIMIT     = 10
)

var columns = []models.GridColumn{
	{
		Typ:      models.String,
		Label:    "Name",
		Key:      "Name",
		Renderer: "name",
	},
	{
		Typ:      models.String,
		Label:    "Status",
		Key:      "Status",
		Renderer: "status",
		Editable: true,
		EditOptions: models.GridEditOptions{
			EditType: models.EditSelect,
			EditProps: models.GridSelectEditProps{
				Id:    "Id",
				Name:  "status",
				Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
				Options: []models.SelectOption{
					{Label: "Active", Value: "active"},
					{Label: "Inactive", Value: "inactive"},
				},
			},
		},
	},
	{
		Typ:      models.String,
		Label:    "Role",
		Key:      "Position",
		Editable: true,
		EditOptions: models.GridEditOptions{
			EditType: models.EditInput,
			EditProps: models.GridInputEditProps{
				Id:    "Id",
				Typ:   models.InputTypeText,
				Name:  "position",
				Class: "mt-1 w-full rounded-md border-gray-200 bg-white text-sm text-gray-700 shadow-sm",
			},
		},
	},
	{
		Typ:      models.Array,
		Label:    "Badges",
		Key:      "Badges",
		Renderer: "badges",
	},
	{
		Typ:   models.String,
		Label: "Email",
		Key:   "Email",
	},
}

var gridCtx = models.GridContext[models.GridColumn]{
	Title:       "Customers",
	Subheading:  "",
	Description: "Example Gird with filters, pagination, export etc...",
	Columns:     columns,
	IdField:     "Id",
}

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
		paginatedData = []models.GridDataRow{}
	}

	var pageOptions = models.GridPagination{
		Current:    page,
		TotalPages: totalPages,
		Limit:      limit,
	}

	return Render(c, http.StatusOK, grid.Grid[models.GridDataRow](gridCtx, paginatedData, pageOptions))
}

func GridFilterHandler(c echo.Context) error {
	page, pageErr := strconv.Atoi(c.QueryParam("page"))
	limit, limitErr := strconv.Atoi(c.QueryParam("limit"))

	if pageErr != nil {
		page = DEFAULT_PAGE_SIZE
	}

	if limitErr != nil {
		limit = DEFAULT_LIMIT
	}

	filterKey := "name"
	filterValue := c.FormValue(filterKey)
	fmt.Println(filterValue)

	var data []models.GridDataRow

	for _, row := range services.Data {
		fieldValue := getFieldStringValue(row, filterKey)
		if strings.Contains(strings.ToLower(fieldValue), strings.ToLower(filterValue)) {
			data = append(data, row)
		}
	}
	totalPages := int(math.Ceil(float64(len(data)) / float64(limit)))
	offset := (page - 1) * limit

	start := offset
	end := offset + limit
	if end > len(data) {
		end = len(data)
	}
	paginatedData := data[start:end]

	var pageOptions = models.GridPagination{
		Current:    page,
		TotalPages: totalPages,
		Limit:      limit,
	}

	return Render(c, http.StatusOK, grid.Grid(gridCtx, paginatedData, pageOptions))
}

func GridRowHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(services.Data, id)
	return Render(c, http.StatusOK, grid.RenderRow(columns, row))
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
	return Render(c, http.StatusOK, grid.RenderRow(columns, updatedRow))
}

func DeleteGridRowHandler(c echo.Context) error {
	id := c.Param("id")
	services.DeleteById(id)
	return c.NoContent(http.StatusOK)
}

func RenderEditGridHandler(c echo.Context) error {
	id := c.Param("id")
	row := services.FilterById(services.Data, id)
	return Render(c, http.StatusOK, grid.EditRow(columns, row))
}

func getFieldStringValue(row models.GridDataRow, fieldName string) string {
	switch fieldName {
	case "name":
		return row.Name
	default:
		return ""
	}
}
