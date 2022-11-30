package models

type ProjectDeprecatedRelationships struct {

	ImportingUser DeprecatedRelationship `json:"importingUser,omitempty"`

	Org DeprecatedRelationship `json:"org,omitempty"`

	Owner DeprecatedRelationship `json:"owner,omitempty"`

	Target DeprecatedRelationship `json:"target,omitempty"`
}
