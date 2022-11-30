package models

type ProjectDeprecated struct {

	Attributes ProjectDeprecatedAttributes `json:"attributes"`

	// The ID.
	Id string `json:"id"`

	Relationships ProjectDeprecatedRelationships `json:"relationships,omitempty"`

	// Content type.
	Type string `json:"type"`
}
