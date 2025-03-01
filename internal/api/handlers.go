package api

import (
	"InsultAPI/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeleteInsult(c echo.Context) error {
	idSTR := c.Param("id")
	id, err := strconv.Atoi(idSTR)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	err = services.DeleteInsult(id)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Insult not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An internal error occurred"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Insult deleted successfully"})
}

func AddInsult(c echo.Context) error {
	insultText := c.FormValue("insult")
	if insultText == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Insult cannot be empty"})
	}

	err := services.AddInsult(insultText)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add insult"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Insult Added"})
}

func GetRandomInsult(c echo.Context) error {
	insult, err := services.GetRandomInsult()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch insult"})
	}
	return c.JSON(http.StatusOK, map[string]string{"insult": insult.Insult})
}

func GetAllInsults(c echo.Context) error {
	insultList, err := services.GetAllInsults()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch insults"})
	}
	return c.JSON(http.StatusOK, insultList)
}

func AddMultipleInsults(c echo.Context) error {
	var insults []string
	if err := c.Bind(&insults); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input format"})
	}

	if len(insults) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No insults provided"})
	}

	err := services.AddMultipleInsults(insults)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add insults"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Insults added successfully"})
}
