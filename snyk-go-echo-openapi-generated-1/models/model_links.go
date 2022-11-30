package models

type Links struct {

	First LinkProperty `json:"first,omitempty"`

	Last LinkProperty `json:"last,omitempty"`

	Next LinkProperty `json:"next,omitempty"`

	Prev LinkProperty `json:"prev,omitempty"`

	Related LinkProperty `json:"related,omitempty"`

	Self LinkProperty `json:"self,omitempty"`
}
