package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// We create a struct to hold the STRUCTure of the data. It includes a nested struct for the properties.
type GetForecastUrl struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Low      float64 `json:"temp_min"`
		High     float64 `json:"temp_max"`
		Humidity float64 `json:"humidity"`
	}
	Wind struct {
		Speed float64 `json:"speed"`
	}
}

func UnmarshalJSON[T any](data []byte) (T, error) {
	var obj T
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return obj, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return obj, nil
}

func GetWeatherData(url string) (WeatherData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherData{}, err
	}

	weatherData, err := UnmarshalJSON[WeatherData](body)
	if err != nil {
		return WeatherData{}, err
	}

	return weatherData, nil
}

func main() {
	urls := make(map[string]string)
	OW_KEY := os.Getenv("OW_KEY")

	urls["Los Angeles"] = "https://api.openweathermap.org/data/2.5/weather?lat=34.0522&lon=-118.2437&units=imperial&appid=" + OW_KEY
	urls["Atlanta"] = "https://api.openweathermap.org/data/2.5/weather?lat=33.6362&lon=-84.4294&units=imperial&appid=" + OW_KEY
	urls["New York"] = "https://api.openweathermap.org/data/2.5/weather?lat=40.7833&lon=-73.9666&units=imperial&appid=" + OW_KEY
	urls["Chicago"] = "https://api.openweathermap.org/data/2.5/weather?lat=41.8376&lon=-87.6818&units=imperial&appid=" + OW_KEY

	for city, url := range urls {
		WeatherData, err := GetWeatherData(url)
		if err != nil {
			fmt.Printf("Failed to get weather data for %s: %v\n", city, err)
			continue
		}

		humidity := WeatherData.Main.Humidity
		temp := WeatherData.Main.Temp
		low := WeatherData.Main.Low
		wind_speed := WeatherData.Wind.Speed

		fmt.Printf("The current humidity for the area of %s is %.2f%%\n", city, humidity)
		fmt.Printf("The current temperature for the area of %s is %d degrees\n", city, int(temp))
		fmt.Printf("The current low temperature for the area of %s is %d degrees\n", city, int(low))
		fmt.Printf("The current wind speed for the area of %s is %.2f mph\n", city, wind_speed)
	}
}
