# Weather Forecast CLI

This is a simple command-line interface (CLI) application written in Go that fetches and prints the current weather data for a list of cities.
It demonstrates fetching and processing JSON data from an API, error handling, and working with maps and loops in Go.

## How it works

The application uses the National Weather Service API to fetch weather data for specific geographical points (latitude, longitude). It prints the humidity, temperature, and dew point for each city.

## Code Structure

The code is organized into two main parts:

1. [`GetWeatherData(url string) (WeatherData, error)`]: This function takes a URL as input, sends a GET request to the National Weather Service API, and returns the weather data for the geographical point represented by the URL.

2. [`FetchURL(url)`]: Takes a URL string as input, performs an HTTP GET request, and returns the response body as a byte slice. 

3. [`UnmarshalJSON[T any](data []byte) (T, error)`]: A generic function that takes a byte slice and unmarshals it into a specified type T. It's used to parse the JSON response into the defined structs

2. [`main()`]: Initializes a map with city names as keys and their corresponding API endpoint URLs as values. It iterates over this map, calls GetWeatherData for each city, and prints the humidity, temperature, and dewpoint for the first period of available data.

## Usage

To run the application, use the `go run` command:

```bash
go run forecast.go
```

## Output

The output will look something like this:

```
The current humidity for the area of Los Angeles is 58%
The current temperature for the area of Los Angeles is 72 F
The current dew point for the area of Los Angeles is -73.33
The current humidity for the area of Atlanta is 36%
The current temperature for the area of Atlanta is 90 F
The current dew point for the area of Atlanta is 15.00
```

## Error Handling

The application includes basic error handling. If it fails to fetch the weather data for a city, it will print an error message and continue with the next city. If the API response does not contain a forecast URL for a city, the application will terminate with an error.

## Future Improvements

Future improvements could include unit tests, support for more cities, and additional weather data such as wind speed and precipitation.