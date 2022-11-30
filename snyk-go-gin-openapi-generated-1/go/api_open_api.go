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

// GetAPIVersion - 
func GetAPIVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ListAPIVersions - 
func ListAPIVersions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
