package models

type DeprecatedRelationship struct {

	Data DeprecatedRelationshipData `json:"data"`

	Links Links `json:"links,omitempty"`

	// Free-form object that may contain non-standard information.
	Meta map[string]interface{} `json:"meta,omitempty"`
}
