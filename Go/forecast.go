package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// We create a struct to hold the STRUCTure of the data. It includes a nested struct for the properties.
type GetForecastUrl struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type WeatherData struct {
	Properties struct {
		Periods []struct {
			RelativeHumidity struct {
				Value int `json:"value"`
			} `json:"relativeHumidity"`
			Temperature int `json:"temperature"`
			Dewpoint    struct {
				Value float64 `json:"value"`
			} `json:"dewpoint"`
		} `json:"periods"`
	} `json:"properties"`
}

func GetWeatherData(url string) (WeatherData, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get forecast: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var getForecastUrl GetForecastUrl
	err = json.Unmarshal(body, &getForecastUrl)
	if err != nil {
		log.Fatalf("Failed to unmarshal forecast: %v", err)
	}

	if getForecastUrl.Properties.Forecast == "" {
		log.Fatal("Forecast URL is empty")
	}

	resp, err = http.Get(getForecastUrl.Properties.Forecast)
	if err != nil {
		log.Fatalf("Failed to get forecast: %v", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatalf("Failed to unmarshal weather data: %v", err)
	}

	if len(weatherData.Properties.Periods) == 0 {
		log.Fatal("No weather data available")
	}
	return weatherData, nil

}

func main() {
	urls := make(map[string]string)

	urls["Los Angeles"] = "https://api.weather.gov/points/34.0522,-118.2437"
	urls["Atlanta"] = "https://api.weather.gov/points/33.6362,-84.4294"
	urls["New York"] = "https://api.weather.gov/points/40.7833,-73.9666"
	urls["Chicago"] = "https://api.weather.gov/points/41.8376,-87.6818"

	for city, url := range urls {
		weatherData, err := GetWeatherData(url)
		if err != nil {
			log.Printf("Failed to get weather data for %s: %v", city, err)
			continue
		}

		humidity := weatherData.Properties.Periods[0].RelativeHumidity.Value
		temp := weatherData.Properties.Periods[0].Temperature
		dewpoint := weatherData.Properties.Periods[0].Dewpoint.Value

		fmt.Printf("The current humidity for the area of %s is %d%%\n", city, humidity)
		fmt.Printf("The current temperature for the area of %s is %d F\n", city, temp)
		fmt.Printf("The current dew point for the area of %s is %.2f\n", city, dewpoint)
	}
}
