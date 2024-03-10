package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/harish876/go-templ-htmx-exploration/models"
	"github.com/harish876/go-templ-htmx-exploration/views/form"
	"github.com/labstack/echo/v4"
)

func FormSubmissionHandler(c echo.Context) error {
	formValues := models.FormValues{
		FirstName:            c.FormValue("first_name"),
		LastName:             c.FormValue("last_name"),
		Email:                c.FormValue("email"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("password_confirmation"),
		MarketingAccept:      c.FormValue("marketing_accept"),
	}

	var formErrors models.FormErrors
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

	return Render(c, http.StatusOK, form.Form(models.FormValues{}, models.FormErrors{}, true, "Success"))
}

func FormHandler(c echo.Context) error {
	return Render(c, http.StatusOK, form.Form(models.FormValues{}, models.FormErrors{}, false, ""))
}
