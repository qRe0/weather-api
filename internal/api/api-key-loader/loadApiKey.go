package api_key_loader

import (
	"encoding/json"
	"os"
	apid "weather-api/internal/api/api-data-struct"
)

func LoadApiKey(filename string) (apid.ApiConfigData, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apid.ApiConfigData{}, err
	}

	var data apid.ApiConfigData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return apid.ApiConfigData{}, err
	}

	return data, nil
}
