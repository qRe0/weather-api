package server_response

import (
	"github.com/labstack/echo/v4"
	"net/http"
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

	return c.JSON(http.StatusOK, data)
}
