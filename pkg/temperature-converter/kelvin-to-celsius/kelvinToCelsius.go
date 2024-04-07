package kelvin_to_celsius

import "math"

func KelvinToCelsius(k float64) float64 {
	temp := math.Round((k-273.15)*10) / 10
	return temp
}
