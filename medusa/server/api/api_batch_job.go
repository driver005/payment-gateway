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

// GetBatchJobs - List Batch Jobs
func GetBatchJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetBatchJobsBatchJob - Get a Batch Job
func GetBatchJobsBatchJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostBatchJobs - Create a Batch Job
func PostBatchJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostBatchJobsBatchJobCancel - Cancel a Batch Job
func PostBatchJobsBatchJobCancel(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// PostBatchJobsBatchJobConfirmProcessing - Confirm a Batch Job
func PostBatchJobsBatchJobConfirmProcessing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
