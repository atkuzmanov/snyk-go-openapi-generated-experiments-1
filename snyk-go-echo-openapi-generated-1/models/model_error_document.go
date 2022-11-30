package models

type ErrorDocument struct {

	Errors []Error `json:"errors"`

	Jsonapi JsonApi `json:"jsonapi"`
}
