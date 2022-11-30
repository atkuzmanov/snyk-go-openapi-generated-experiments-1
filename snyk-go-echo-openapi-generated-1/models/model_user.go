package models

type User struct {

	Attributes UserAttributes `json:"attributes"`

	// The Snyk ID corresponding to this user
	Id string `json:"id"`

	// Content type.
	Type string `json:"type"`
}
