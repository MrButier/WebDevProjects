package main

import (
	"InsultAPI/internal/api"
	"InsultAPI/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
	"log"
	"strconv"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, using environment variables instead.")
	}

	viper.SetDefault("server_port", 1323)
	viper.SetDefault("rate_limit", 5)

	port := viper.GetInt("server_port")
	rateLimit := rate.Limit(viper.GetFloat64("rate_limit"))

	portStr := ":" + strconv.Itoa(port)

	database.InitDB()
	defer database.CloseDB()

	e := echo.New()

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rateLimit)))

	api.SetupRouter(e)

	log.Printf("Server running on port %d...", port)
	e.Logger.Fatal(e.Start(portStr))
}
