package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var appName string
var apiKey = "4b1e3c9fc55aab7d840a37117f2c9bd0" // Replace with your actual API key

func main() {
	log.SetOutput(os.Stderr)
	appName = "dynamicWelcomeService"
	if os.Getenv("APP_NAME") != "" {
		appName = os.Getenv("APP_NAME")
	}

	log.Println("Starting the dynamicWelcomeService server...")
	router := gin.New()
	router.GET("/", homeHandler)
	router.GET("/info", infoHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}

func homeHandler(ctx *gin.Context) {
	clientIP := ctx.ClientIP()
	uuid := uuid.New()

	weather, err := getWeatherForAthens()
	if err != nil {
		log.Println("Failed to get weather:", err)
		ctx.String(http.StatusInternalServerError, "Error fetching weather information\n")
		return
	}

	country := "Greece" // Since we're specifically fetching weather for Athens, Greece
	message := "Welcome to the dynamicWelcomeService demo"
	if os.Getenv("WELCOME_MESSAGE") != "" {
		message = os.Getenv("WELCOME_MESSAGE")
	}

	response := fmt.Sprintf("%s\nYour IP: %s\nCountry: %s\nSession ID: %s\nWeather in Athens: %s\n", message, clientIP, country, uuid, weather)
	ctx.String(http.StatusOK, response)
}

func getWeatherForAthens() (string, error) {
	response, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=Athens,gr&appid=%s&units=metric", apiKey))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return "", err
	}

	// Extract weather information. Simplified for demonstration.
	if weather, found := data["weather"].([]interface{}); found && len(weather) > 0 {
		if description, ok := weather[0].(map[string]interface{})["description"].(string); ok {
			return description, nil
		}
	}

	return "Not available", nil
}

func infoHandler(ctx *gin.Context) {
	appVersion := os.Getenv("VERSION")
	infoMessage := fmt.Sprintf("%s: A simple demo application.", appName)
	if len(appVersion) > 0 {
		infoMessage = fmt.Sprintf("%s version %s", infoMessage, appVersion)
	}
	ctx.String(http.StatusOK, "%s\n", infoMessage)
}
