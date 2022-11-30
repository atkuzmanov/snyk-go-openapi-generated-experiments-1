package models

type LinkProperty struct {

	// A string containing the link’s URL.
	Href string `json:"href"`

	// Free-form object that may contain non-standard information.
	Meta map[string]interface{} `json:"meta,omitempty"`
}
