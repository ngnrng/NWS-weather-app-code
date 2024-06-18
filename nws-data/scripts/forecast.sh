#!/bin/bash

# Get forecast urls
urls=$(curl -s "https://api.weather.gov/points/34.0522,-118.2437" | jq -r '.properties.forecast')

# Check if forecast URL is empty
if [ -z "$urls" ]; then
    echo "Failed to retrieve forecast URL. Exiting."
    exit 1
fi

# Get forecast data
forecast_data=$(curl -s "$urls")

# Check if forecast data is empty
if [ -z "$forecast_data" ]; then
    echo "Failed to retrieve forecast data. Exiting."
    exit 1
fi

# Get current relative humidity and parse json response to extract values
humidity=$(echo "$forecast_data" | jq -r '.properties.periods[0].relativeHumidity.value')

# Get current temperature and parse json response to extract values
temp=$(echo "$forecast_data" | jq -r '.properties.periods[0].temperature')

# Get current dew point and parse json response to extract values
dewpoint=$(echo "$forecast_data"  | jq -r '.properties.periods[0].dewpoint.value')

# Check if any of the values are empty
if [ -z "$humidity" ] || [ -z "$temp" ] || [ -z "$dewpoint" ]; then
    echo "Failed to retrieve one or more weather parameters. Exiting."
    exit 1
fi

# Echo results
echo "The current humidity for the Los Angeles area is $humidity%"
echo "The current temperature for the Los Angeles area is $temp F"
echo "The current dew point for the Los Angeles area is $dewpoint"

