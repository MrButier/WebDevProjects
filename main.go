package main

import (
	"InsultAPI/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func main() {
	database.InitDB()
	e := echo.New()

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(5)))

	e.GET("/insult", getRandomInsult)
	e.POST("/insult", addInsult)

	e.GET("/insults/all", getAllInsults)
	e.POST("/insults", addMultipleInsults)

	e.DELETE("/insult/:id", deleteInsult)

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))

	defer database.CloseDB()
}

func deleteInsult(c echo.Context) error {
	idSTR := c.Param("id")
	id, err := strconv.Atoi(idSTR)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid ID format"})
	}

	err = database.DeleteInsult(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to delete insult"})
	}

	return c.JSON(200, map[string]string{"message": "Insult Deleted"})
}

func addInsult(c echo.Context) error {
	insultText := c.FormValue("insult")

	if insultText == "" {
		return c.JSON(400, map[string]string{"error": "Insult cannot be empty"})
	}

	err := database.AddInsult(insultText)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to add insult"})
	}

	return c.JSON(200, map[string]string{"message": "Insult Added"})
}

func getRandomInsult(c echo.Context) error {
	insult, _ := database.GetRandomInsult()
	return c.JSON(200, map[string]string{"insult": insult.Insult})
}

func getAllInsults(c echo.Context) error {
	insultList, _ := database.GetAllInsults()
	return c.JSON(200, insultList)
}

func addMultipleInsults(c echo.Context) error {
	var insults []string

	// Bind the request body to a slice of strings
	if err := c.Bind(&insults); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input format"})
	}

	if len(insults) == 0 {
		return c.JSON(400, map[string]string{"error": "No insults provided"})
	}

	// Loop through and add each insult
	for _, insult := range insults {
		err := database.AddInsult(insult)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to add insults"})
		}
	}

	return c.JSON(200, map[string]string{"message": fmt.Sprintf("%d insults added successfully!", len(insults))})
}
