package models

type GetOrgsTarget200Response struct {

	Data Target `json:"data"`

	Jsonapi JsonApi `json:"jsonapi"`

	Links Links `json:"links,omitempty"`
}
