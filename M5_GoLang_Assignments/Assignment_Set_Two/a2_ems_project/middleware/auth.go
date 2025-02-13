package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddlewareGin(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Missing or invalid Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Decode the Base64-encoded credentials
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			fmt.Println("Failed to decode Authorization header:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		// Split the name and password
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			fmt.Println("Invalid credentials format:", string(payload))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]
		fmt.Printf("Received name: %s, password: %s\n", username, password)

		// Validate credentials against the database
		var storedPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&storedPassword)
		if err != nil {
			fmt.Println("User not found or error querying database:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if storedPassword != password {
			fmt.Println("Password mismatch")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		fmt.Println("Authentication successful")
		c.Next()
	}
}
