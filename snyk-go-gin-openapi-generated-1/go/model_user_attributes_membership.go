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
	"time"
)

type UserAttributesMembership struct {

	// The date the membership was established.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Whether the membership is a direct, or indirect membership.
	Strategy string `json:"strategy,omitempty"`
}