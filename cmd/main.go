package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	h "weather-api/internal/middleware/server-response"
)

var port = 8080

func main() {
	e := echo.New()
	e.Static("/static", "client-part")

	e.GET("/weather/:city", h.ResponseHandler)

	err := e.Start(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error starting server on port %d: %s", port, err)
	}

}
