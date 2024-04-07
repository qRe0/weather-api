package weather_struct

type Weather struct {
	Description string `json:"description"`
}

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
	Weather []Weather `json:"weather"`
}
