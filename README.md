# Weather Forecast CLI

This is a simple command-line interface (CLI) application written in Go that fetches and prints the current weather data for a list of cities.
It demonstrates fetching and processing JSON data from an API, error handling, and working with maps and loops in Go.

## How it works

The application now uses the OpenWeatherMap API to fetch weather data for specific cities. It prints the humidity, temperature, low temperature, high temperature, and wind speed for each city.

## Code Structure

The code is organized into three main parts:

1. `GetWeatherData(url string) (WeatherData, error)`: This function takes a URL as input, sends a GET request to the OpenWeatherMap API, and returns the weather data for the city represented by the URL.

2. `UnmarshalJSON[T any](data []byte) (T, error)`: A generic function that takes a byte slice and unmarshals it into a specified type T. It's used to parse the JSON response into the defined structs.

3. `main()`: Initializes a map with city names as keys and their corresponding API endpoint URLs as values. It iterates over this map, calls `GetWeatherData` for each city, and prints the humidity, temperature, low temperature, high temperature, and wind speed for each city.

## Usage

To run the application, use the `go run` command:

```go run forecast.go```

## Output

The output will look something like this:

```The current humidity for the area of Los Angeles is 58%
The current temperature for the area of Los Angeles is 72 degrees
The current low temperature for the area of Los Angeles is 68 degrees
The current high temperature for the area of Los Angeles is 76 degrees
The current wind speed for the area of Los Angeles is 5.10 mph
...
```

### Error Handling

The application includes basic error handling. If it fails to fetch the weather data for a city, it will print an error message and continue with the next city. If the API response does not contain the expected weather data for a city, the application will terminate with an error.

### Future Improvements

Future improvements could include unit tests, support for more cities, and additional weather data such as precipitation and visibility.