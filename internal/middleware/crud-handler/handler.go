package crud_handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	srp "weather-api/internal/middleware/server-response"
)

func MenuHandler(c echo.Context) error {
	return c.File("client-part/menu.html")
}

func WeatherHandler(c echo.Context) error {
	err := srp.ResponseHandler(c)
	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "")
}

func GetWeatherHandler(c echo.Context) error {
	city := c.FormValue("city")
	return c.Redirect(http.StatusMovedPermanently, "/weather/"+city)
}
