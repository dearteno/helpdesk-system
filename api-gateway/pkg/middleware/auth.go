package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// SupabaseAuth middleware to validate JWT tokens from Supabase
func SupabaseAuth(supabaseURL, supabaseKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for public routes
		if isPublicRoute(c.FullPath()) {
			c.Next()
			return
		}

		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		// Check for Bearer token
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header must start with Bearer",
			})
			return
		}

		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify token with Supabase
		user, err := verifyToken(token, supabaseURL, supabaseKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Add user info to the context
		c.Set("user", user)
		c.Next()
	}
}

// ProxyService forwards requests to the appropriate gRPC service
func ProxyService(serviceURL, grpcMethod string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// This is a simplified version. In a real implementation, you would:
		// 1. Convert the REST request to gRPC
		// 2. Make the gRPC call to the service
		// 3. Convert the gRPC response back to REST
		// 4. Return the response to the client

		// For now, we'll just return a placeholder
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Proxied to %s at %s", grpcMethod, serviceURL),
			// In a real implementation, this would be the actual response from the service
		})
	}
}

// Helper functions

func isPublicRoute(path string) bool {
	publicRoutes := []string{
		"/health",
		"/api/v1/faq", // Making FAQs public
		"/api/v1/auth",
	}

	for _, route := range publicRoutes {
		if strings.HasPrefix(path, route) {
			return true
		}
	}

	return false
}

type SupabaseUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func verifyToken(token, supabaseURL, supabaseKey string) (*SupabaseUser, error) {
	// In a real implementation, you would make an HTTP request to Supabase Auth API
	// to validate the token and get user information

	if token == "" {
		return nil, errors.New("invalid token")
	}

	// Make an HTTP request to Supabase Auth API to verify the token
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", supabaseURL+"/auth/v1/user", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add the necessary headers
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("apikey", supabaseKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid or expired token")
	}

	// Parse the user data from the response
	var userData SupabaseUser
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return nil, fmt.Errorf("failed to parse user data: %w", err)
	}

	return &userData, nil
}
