package weather_query

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	load "weather-api/internal/api/api-key-loader"
	ws "weather-api/internal/weather/weather-struct"
)

func Query(city string) (ws.WeatherData, error) {
	ApiConfig, err := load.LoadApiKey("internal/api/.apiConfig")
	if err != nil {
		return ws.WeatherData{}, err
	}

	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?APPID=" + ApiConfig.ApiKey + "&q=" + city)
	if err != nil {
		return ws.WeatherData{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %s\n", err)
		}
	}(resp.Body)

	var data ws.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ws.WeatherData{}, err
	}

	return data, nil
}
