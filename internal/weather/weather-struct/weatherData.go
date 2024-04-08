package weather_struct

type Weather struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherData struct {
	City string `json:"name"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
	Weather []Weather `json:"weather"`
}
