/*
 * Snyk API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: REST
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ErrorDocument struct {

	Errors []ModelError `json:"errors"`

	Jsonapi *JsonApi `json:"jsonapi"`
}