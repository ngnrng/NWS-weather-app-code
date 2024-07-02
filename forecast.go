package main

import (
	"bytes"
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

type CityMetrics struct {
	City      string
	Humidity  float64
	Temp      int
	Low       int
	High      int
	WindSpeed float64
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

func sendMetrics(citiesMetrics []CityMetrics) error {
	var data string
	for _, cm := range citiesMetrics {
		data += fmt.Sprintf(
			"relative_humidity{city=\"%s\"} %.2f\n"+
				"current_temperature{city=\"%s\"} %d\n"+
				"low_temp{city=\"%s\"} %d\n"+
				"high_temp{city=\"%s\"} %d\n"+
				"wind_speed{city=\"%s\"} %.2f\n",
			cm.City, cm.Humidity, cm.City, cm.Temp, cm.City, cm.Low, cm.City, cm.High, cm.City, cm.WindSpeed,
		)
	}
	// Print metrics
	fmt.Println("Metrics being sent:\n", data)

	// Endpoint where the metrics will be sent
	url := "http://pushgateway.monitoring.svc.cluster.local:9091/metrics/job/weather_metrics"

	// Create a new POST request with the formatted data
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/plain")

	// Send the request using the http client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the POST request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send metrics, status code: %d", resp.StatusCode)
	}

	fmt.Println("Metrics successfully sent")
	return nil
}

func main() {
	urls := make(map[string]string)
	OW_KEY := os.Getenv("OW_KEY")

	urls["Los Angeles"] = "https://api.openweathermap.org/data/2.5/weather?lat=34.0522&lon=-118.2437&units=imperial&appid=" + OW_KEY
	urls["Atlanta"] = "https://api.openweathermap.org/data/2.5/weather?lat=33.6362&lon=-84.4294&units=imperial&appid=" + OW_KEY
	urls["New York"] = "https://api.openweathermap.org/data/2.5/weather?lat=40.7833&lon=-73.9666&units=imperial&appid=" + OW_KEY
	urls["Chicago"] = "https://api.openweathermap.org/data/2.5/weather?lat=41.8376&lon=-87.6818&units=imperial&appid=" + OW_KEY

	var citiesMetrics []CityMetrics

	for city, url := range urls {
		WeatherData, err := GetWeatherData(url)
		if err != nil {
			fmt.Printf("Failed to get weather data for %s: %v\n", city, err)
			continue
		}

		cm := CityMetrics{
			City:      city,
			Humidity:  WeatherData.Main.Humidity,
			Temp:      int(WeatherData.Main.Temp),
			Low:       int(WeatherData.Main.Low),
			High:      int(WeatherData.Main.High),
			WindSpeed: WeatherData.Wind.Speed,
		}

		citiesMetrics = append(citiesMetrics, cm)

		fmt.Printf("The current humidity for the area of %s is %.2f%%\n", city, cm.Humidity)
		fmt.Printf("The current temperature for the area of %s is %d degrees\n", city, cm.Temp)
		fmt.Printf("The current low temperature for the area of %s is %d degrees\n", city, cm.Low)
		fmt.Printf("The current high temperature for the area of %s is %d degrees\n", city, cm.High)
		fmt.Printf("The current wind speed for the area of %s is %.2f mph\n", city, cm.WindSpeed)
	}

	err := sendMetrics(citiesMetrics)
	if err != nil {
		fmt.Printf("Failed to send metrics: %v\n", err)
	}
}
