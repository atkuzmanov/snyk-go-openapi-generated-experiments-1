package models

type ErrorSource struct {

	// A string indicating which URI query parameter caused the error.
	Parameter string `json:"parameter,omitempty"`

	// A JSON Pointer [RFC6901] to the associated entity in the request document.
	Pointer string `json:"pointer,omitempty"`
}
