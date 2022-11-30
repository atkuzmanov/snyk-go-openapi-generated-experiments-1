package models

// ErrorLink - A link that leads to further details about this particular occurrance of the problem.
type ErrorLink struct {

	About LinkProperty `json:"about,omitempty"`
}
