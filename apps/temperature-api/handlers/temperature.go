package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TemperatureHandler handles temperature requests
type TemperatureHandler struct {
}

// NewTemperatureHandler creates a new TemperatureHandler
func NewTemperatureHandler() *TemperatureHandler {
	return &TemperatureHandler{}
}

// RegisterRoutes registers the temperature routes
func (h *TemperatureHandler) RegisterRoutes(router *gin.RouterGroup) {
	temperature := router.Group("/temperature")
	{
		temperature.GET("", h.GetTemperature)
		temperature.GET("/:sensor_id", h.GetTemperatureBySensorID)
	}
}

// GetTemperature handles GET /temperature
func (h *TemperatureHandler) GetTemperature(c *gin.Context) {
	var sensorID string

	location := c.Query("location")

	// No sensor ID is provided, generate one based on location
	switch location {
	case "Living Room":
		sensorID = "1"
	case "Bedroom":
		sensorID = "2"
	case "Kitchen":
		sensorID = "3"
	default:
		location = "Unknown"
		sensorID = "0"
	}

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"description": "random value",
		"location":    location,
		"sensor_id":   sensorID,
		"sensor_type": "temperature",
		"status":      "active",
		"timestamp":   time.Now(),
		"value":       rand.Float64() * 30,
		"unit":        "°C",
	})
}

// GetTemperatureBySensorID handles GET /temperature/:sensor_id
func (h *TemperatureHandler) GetTemperatureBySensorID(c *gin.Context) {
	var location string

	sensorID := c.Param("sensor_id")

	// No location is provided, use a default based on sensor ID
	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		sensorID = "0"
		location = "Unknown"
	}

	// Return the temperature data
	c.JSON(http.StatusOK, gin.H{
		"description": "random value",
		"location":    location,
		"sensor_id":   sensorID,
		"sensor_type": "temperature",
		"status":      "active",
		"timestamp":   time.Now(),
		"value":       rand.Float64() * 30,
		"unit":        "°C",
	})
}
