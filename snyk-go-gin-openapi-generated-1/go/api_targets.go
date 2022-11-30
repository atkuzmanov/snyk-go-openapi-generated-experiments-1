/*
 * Snyk API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: REST
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteOrgsTarget - Delete target by target ID
func DeleteOrgsTarget(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetOrgsTarget - Get target by org ID
func GetOrgsTarget(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetOrgsTargets - Get targets by org ID
func GetOrgsTargets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
