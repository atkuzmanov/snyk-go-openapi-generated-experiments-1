/*
 * Snyk API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: REST
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type HelloWorldAttributes struct {

	BetaField string `json:"betaField"`

	Message string `json:"message"`

	RequestSubject HelloWorldAttributesRequestSubject `json:"requestSubject"`
}
