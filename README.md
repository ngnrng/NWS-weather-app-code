This bash script retrieves current weather information for the Los Angeles, Atlanta, and NYC areas using data from the National Weather Service API.

## Steps the Script Goes Through

1. **Get Forecast URLs**: 
   - The script sends a request to the National Weather Service API to obtain the forecast URLS for the 3 areas.
   - It extracts the forecast URL from the JSON response using `jq`.

2. **Check Forecast URL**:
   - Error Handling: It checks if the forecast URL is empty. If empty, it displays an error message and exits.

3. **Get Forecast Data**:
   - The script sends a request to the forecast URL obtained in the previous step to fetch the forecast data.
   - It stores the forecast data in a variable.

4. **Check Forecast Data**:
   - Error Handling: It checks if the forecast data is empty. If empty, it displays an error message and exits.

5. **Parse Forecast Data**:
   - The script parses the JSON response to extract the current relative humidity, temperature, and dew point for the 3 areas.

6. **Check Weather Parameters**:
   - Error Handling: It checks if any of the extracted weather parameters (humidity, temperature, dew point) are empty. If any parameter is empty, it displays an error message and exits.

7. **Display Results**:
   - Finally, the script echoes the current humidity, temperature, and dew point for the 3 areas.