package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"strconv"

	h "weather-api/internal/middleware/crud-handler"
)

var port = 8080

func main() {
	e := echo.New()
	e.Static("/static", "client-part")

	e.GET("/menu", h.MenuHandler)
	e.GET("/weather/:city", h.WeatherHandler)
	e.POST("/get-weather", h.GetWeatherHandler)

	err := e.Start(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error launching server on port %d: %s", port, err)
	}
}
