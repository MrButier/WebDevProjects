package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(e *echo.Echo) {
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(5)))

	e.GET("/insult", GetRandomInsult)
	e.POST("/insult", AddInsult)
	e.GET("/insults", GetAllInsults)
	e.POST("/insults", AddMultipleInsults)
	e.DELETE("/insult/:id", DeleteInsult)
}
