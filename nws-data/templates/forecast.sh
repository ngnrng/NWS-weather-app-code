#!/bin/bash

{{- range .Values.cities }}
# Define city name and coordinates
city="{{ .name }}"
coordinates="{{ .coordinates }}"

# Get forecast urls for the specified city
urls=$(curl -s "https://api.weather.gov/points/$coordinates" | jq -r '.properties.forecast')

# Check if forecast URL is empty
if [ -z "$urls" ]; then
    echo "Failed to retrieve forecast URL for $city. Skipping."
    continue
fi

# Get forecast data
forecast_data=$(curl -s "$urls")

# Check if forecast data is empty
if [ -z "$forecast_data" ]; then
    echo "Failed to retrieve forecast data for $city. Skipping."
    continue
fi

# Get current relative humidity and parse json response to extract values
humidity=$(echo "$forecast_data" | jq -r '.properties.periods[0].relativeHumidity.value')

# Get current temperature and parse json response to extract values
temp=$(echo "$forecast_data" | jq -r '.properties.periods[0].temperature')

# Get current dew point and parse json response to extract values
dewpoint=$(echo "$forecast_data"  | jq -r '.properties.periods[0].dewpoint.value')

# Check if any of the values are empty
if [ -z "$humidity" ] || [ -z "$temp" ] || [ -z "$dewpoint" ]; then
    echo "Failed to retrieve one or more weather parameters for $city. Skipping."
    continue
fi

# Echo results
echo "The current humidity for the $city area is $humidity%"
echo "The current temperature for the $city area is $temp F"
echo "The current dew point for the $city area is $dewpoint"
echo
{{- end }}

