package models

type Target struct {

	Attributes TargetAttributes `json:"attributes"`

	// The Snyk ID corresponding to this target
	Id string `json:"id"`

	Relationships TargetRelationships `json:"relationships,omitempty"`

	// Content type.
	Type string `json:"type"`
}
