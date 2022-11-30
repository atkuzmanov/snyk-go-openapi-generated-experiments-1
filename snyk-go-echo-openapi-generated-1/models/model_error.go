package models

type Error struct {

	// An application-specific error code, expressed as a string value.
	Code string `json:"code,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail"`

	// A unique identifier for this particular occurrence of the problem.
	Id string `json:"id,omitempty"`

	Links ErrorLink `json:"links,omitempty"`

	Meta map[string]interface{} `json:"meta,omitempty"`

	Source ErrorSource `json:"source,omitempty"`

	// The HTTP status code applicable to this problem, expressed as a string value.
	Status string `json:"status"`

	// A short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title string `json:"title,omitempty"`
}
