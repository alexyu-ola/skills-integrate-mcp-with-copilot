// Weather API Service
//
// A RESTful web API built with Go Gin framework that provides weather
// information for a given city.
package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// WeatherInfo holds weather data for a city.
type WeatherInfo struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
	Condition   string  `json:"condition"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	WindUnit    string  `json:"wind_unit"`
}

// weatherDatabase is a simple in-memory store of weather data for cities.
var weatherDatabase = map[string]WeatherInfo{
	"beijing": {
		City:        "Beijing",
		Temperature: 12.0,
		Unit:        "°C",
		Condition:   "Partly Cloudy",
		Humidity:    55,
		WindSpeed:   15.0,
		WindUnit:    "km/h",
	},
	"shanghai": {
		City:        "Shanghai",
		Temperature: 18.5,
		Unit:        "°C",
		Condition:   "Sunny",
		Humidity:    60,
		WindSpeed:   10.0,
		WindUnit:    "km/h",
	},
	"guangzhou": {
		City:        "Guangzhou",
		Temperature: 26.0,
		Unit:        "°C",
		Condition:   "Cloudy",
		Humidity:    75,
		WindSpeed:   8.0,
		WindUnit:    "km/h",
	},
	"shenzhen": {
		City:        "Shenzhen",
		Temperature: 25.5,
		Unit:        "°C",
		Condition:   "Light Rain",
		Humidity:    80,
		WindSpeed:   12.0,
		WindUnit:    "km/h",
	},
	"chengdu": {
		City:        "Chengdu",
		Temperature: 15.0,
		Unit:        "°C",
		Condition:   "Foggy",
		Humidity:    85,
		WindSpeed:   5.0,
		WindUnit:    "km/h",
	},
	"hangzhou": {
		City:        "Hangzhou",
		Temperature: 17.0,
		Unit:        "°C",
		Condition:   "Sunny",
		Humidity:    58,
		WindSpeed:   9.0,
		WindUnit:    "km/h",
	},
	"wuhan": {
		City:        "Wuhan",
		Temperature: 16.0,
		Unit:        "°C",
		Condition:   "Overcast",
		Humidity:    70,
		WindSpeed:   11.0,
		WindUnit:    "km/h",
	},
	"new york": {
		City:        "New York",
		Temperature: 8.0,
		Unit:        "°C",
		Condition:   "Windy",
		Humidity:    45,
		WindSpeed:   25.0,
		WindUnit:    "km/h",
	},
	"london": {
		City:        "London",
		Temperature: 10.0,
		Unit:        "°C",
		Condition:   "Rainy",
		Humidity:    78,
		WindSpeed:   20.0,
		WindUnit:    "km/h",
	},
	"tokyo": {
		City:        "Tokyo",
		Temperature: 14.0,
		Unit:        "°C",
		Condition:   "Clear",
		Humidity:    50,
		WindSpeed:   7.0,
		WindUnit:    "km/h",
	},
	"paris": {
		City:        "Paris",
		Temperature: 11.0,
		Unit:        "°C",
		Condition:   "Cloudy",
		Humidity:    65,
		WindSpeed:   18.0,
		WindUnit:    "km/h",
	},
	"sydney": {
		City:        "Sydney",
		Temperature: 22.0,
		Unit:        "°C",
		Condition:   "Sunny",
		Humidity:    55,
		WindSpeed:   14.0,
		WindUnit:    "km/h",
	},
}

// setupRouter creates and configures the Gin router.
func setupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/weather", getWeather)
	}

	return r
}

// getWeather handles GET /api/v1/weather?city=<city>
// It returns weather information for the specified city.
func getWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "city parameter is required",
		})
		return
	}

	key := normalizeCity(city)
	info, found := weatherDatabase[key]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "weather data not found for city: " + city,
		})
		return
	}

	c.JSON(http.StatusOK, info)
}

// normalizeCity converts a city name to lowercase for case-insensitive lookup.
func normalizeCity(city string) string {
	return strings.ToLower(city)
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
