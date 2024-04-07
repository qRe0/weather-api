package server_response

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	q "weather-api/internal/middleware/weather-query"
	conv "weather-api/pkg/temperature-converter/kelvin-to-celsius"
)

func ResponseHandler(c echo.Context) error {
	city := c.Param("city")
	data, err := q.Query(city)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	data.Main.Temperature = conv.KelvinToCelsius(data.Main.Temperature)

	baseFile, err := os.ReadFile("client-part/index.html")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	modFile := strings.ReplaceAll(string(baseFile), "city", data.City)
	modFile = strings.ReplaceAll(modFile, "temp", fmt.Sprintf("%.2f", data.Main.Temperature))
	modFile = strings.ReplaceAll(modFile, "descr", data.Weather[0].Description)

	modFile = strings.ReplaceAll(modFile, "bg.jpeg", "/static/bg.jpeg")

	return c.HTML(http.StatusOK, modFile)
}
