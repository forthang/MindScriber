package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var messages []string // Slice to store messages

func main() {
	router := gin.Default()

	// CORS setup
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")              // Allow all origins
		c.Header("Access-Control-Allow-Headers", "Content-Type")  // Allow specific headers
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS") // Allow specific methods

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // Respond with 204 No Content for OPTIONS requests
			return
		}

		c.Next()
	})

	// Endpoint to handle incoming messages
	router.POST("/messages", func(c *gin.Context) {
		var json struct {
			Message string `json:"message"`
		}

		// Bind JSON to struct
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Store the message
		messages = append(messages, json.Message)

		// Respond with the saved message
		c.JSON(http.StatusCreated, gin.H{"message": json.Message})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
