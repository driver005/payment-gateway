/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAuth - User Logout
func DeleteAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAuth - Get Current User
func GetAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostAuth - User Login
func PostAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
