package models

type GetOrgsTargets200Response struct {

	Data []Target `json:"data"`

	Jsonapi JsonApi `json:"jsonapi"`

	Links Links `json:"links"`
}
